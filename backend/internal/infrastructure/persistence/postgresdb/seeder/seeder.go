package postgresdb

import (
	"fmt"
	"math/rand/v2"
	"os"
	"time"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Entry point for seeding
func Seed(db *gorm.DB) error {
	if err := seedTagTypesAndCategories(db); err != nil {
		return err
	}
	if err := seedCampuses(db); err != nil {
		return err
	}
	if err := seedUsers(db); err != nil {
		return err
	}
	if err := seedEvents(db); err != nil {
		return err
	}

	zap.L().Sugar().Info("Database seeded successfully!")
	return nil
}

func seedTagTypesAndCategories(db *gorm.DB) error {
	tagData := []struct {
		TypeName string
		Tags     []string
	}{
		{"Music", []string{"Concerts", "Music Festivals", "Live Music", "Music Workshops", "DJ Nights"}},
		{"Arts & Culture", []string{"Art Exhibitions", "Theater Performances", "Film Screenings", "Cultural Festivals", "Dance Performances"}},
		{"Food & Drink", []string{"Food Festivals", "Wine Tastings", "Cooking Classes", "Food Tours", "Pop-Up Restaurants", "Beer Festivals"}},
		{"Sports & Fitness", []string{"Sports Events", "Fitness Workshops", "Outdoor Activities", "Marathons"}},
		{"Business & Networking", []string{"Networking Events", "Business Conferences", "Workshops"}},
		{"Family & Kids", []string{"Family-Friendly Events", "Kids Activities", "Parenting Workshops", "Outdoor Adventures", "Educational Events"}},
		{"Technology", []string{"Tech Conferences", "Startup Events", "Hackathons", "Gadget Expos", "Innovation Workshops"}},
		{"Comedy & Entertainment", []string{"Comedy Shows", "Stand-Up Comedy", "Improv Nights", "Magic Shows"}},
		{"Charity & Causes", []string{"Charity Galas", "Fundraising Events", "Volunteer Opportunities", "Benefit Concerts"}},
		{"Health & Wellness", []string{"Health Fairs", "Wellness Retreats", "Yoga Classes", "Meditation Workshops"}},
		{"Travel & Adventure", []string{"Travel Expos", "City Tours", "Cultural Experiences", "Cruise Vaccations"}},
		{"Education & Learning", []string{"Teaching Workshops", "Seminars", "Lectures", "Educational Conferences"}},
		{"Fashion & Beauty", []string{"Fashion Shows", "Beauty Workshops", "Makeup Classes", "Style Consultations"}},
	}

	var allCategories []model.Category
	for _, item := range tagData {
		tagType := model.CategoryType{
			Id:   uuid.New(),
			Name: item.TypeName,
		}
		if err := db.Create(&tagType).Error; err != nil {
			return err
		}

		for _, tagName := range item.Tags {
			allCategories = append(allCategories, model.Category{
				Id:             uuid.New(),
				Name:           tagName,
				CategoryTypeId: tagType.Id,
			})
		}
	}

	if len(allCategories) > 0 {
		return db.CreateInBatches(&allCategories, 50).Error
	}
	return nil
}

func seedCampuses(db *gorm.DB) error {
	fake := faker.New()
	rng := rand.New(rand.NewPCG(0, 0))

	campusLogo := fmt.Sprintf("%s/storage/v1/object/public/%s/assets/binus.png",
		os.Getenv("SUPABASE_PROJECT_URL"), os.Getenv("SUPABASE_BUCKET_NAME"))

	initialCampus := model.Campus{
		Id:          uuid.New(),
		Name:        "Binus University",
		LogoUrl:     campusLogo,
		Description: "Best University of The West",
		InviteCode:  "AAAAAA",
	}
	if err := db.Create(&initialCampus).Error; err != nil {
		zap.L().Sugar().Errorf("Failed to seed main campus: %v", err)
		return err
	}

	var campuses []model.Campus
	for i := 0; i < 10; i++ {
		campuses = append(campuses, model.Campus{
			Id:          uuid.New(),
			Name:        fake.Lorem().Sentence(2),
			LogoUrl:     fmt.Sprintf("https://picsum.photos/id/%d/500/350", rng.IntN(200)),
			Description: fake.Lorem().Paragraph(2),
			InviteCode:  fake.RandomLetter() + fake.RandomLetter() + fake.RandomLetter() + fake.RandomLetter() + fake.RandomLetter() + fake.RandomLetter(),
		})
	}

	return db.CreateInBatches(&campuses, 10).Error
}

func seedUsers(db *gorm.DB) error {
	fake := faker.New()
	rng := rand.New(rand.NewPCG(0, 0))

	var campusIds []uuid.UUID
	if err := db.Model(&model.Campus{}).Pluck("id", &campusIds).Error; err != nil {
		return err
	}

	var users []model.User
	for i := 0; i < 50; i++ {
		campusId := campusIds[rng.IntN(len(campusIds))]
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte("testing"), bcrypt.DefaultCost)
		if err != nil {
			zap.L().Sugar().Errorf("Failed to hash password: %v", err)	
			panic(err)
		}

		users = append(users, model.User{
			Id:        uuid.NewString(),
			Provider:  "credential",
			Name:      fake.Person().Name(),
			Email:     fake.Internet().Email(),
			Password:  string(hashedPassword),
			CampusId:  &campusId,
			AvatarUrl: fmt.Sprintf("https://picsum.photos/id/%d/200/200", rng.IntN(200)),
		})
	}

	return db.CreateInBatches(&users, 20).Error
}

func seedEvents(db *gorm.DB) error {
	fake := faker.New()
	rng := rand.New(rand.NewPCG(0, 0))

	var userIds []string
	var campusIds []uuid.UUID
	var categories []model.Category

	if err := db.Model(&model.User{}).Pluck("id", &userIds).Error; err != nil {
		return err
	}
	if err := db.Model(&model.Campus{}).Pluck("id", &campusIds).Error; err != nil {
		return err
	}
	if err := db.Find(&categories).Error; err != nil {
		return err
	}

	for i := 0; i < 500; i++ {
		startTime := fake.Time().TimeBetween(time.Now(), time.Now().AddDate(0, 2, 30))
		var endTime *time.Time
		if rng.IntN(2) == 0 {
			t := fake.Time().TimeBetween(startTime, startTime.Add(2*time.Hour))
			endTime = &t
		}

		categoryCount := rng.IntN(3) + 1
		categoriesToAdd := getUniqueRandomCategories(categories, categoryCount, rng)

		ticketType := []string{"free", "ticketed"}[rng.IntN(2)]

		var tickets []model.TicketType
		if ticketType == "ticketed" {
			ticketTypeCount := rng.IntN(3) + 1
			for j := 0; j < ticketTypeCount; j++ {
				ticket := model.TicketType{
					Id:          uuid.New(),
					Name:        fake.Lorem().Sentence(2),
					Price:       float64(rng.IntN(100) + 1),
					Quantity:    int(rng.IntN(100) + 20), 
				}
				tickets = append(tickets, ticket)
			}
		}

		event := model.Event{
			Id:          uuid.New(),
			Title:       fake.Lorem().Sentence(4),
			Description: ptr(fake.Lorem().Paragraph(3)),
			StartTime:   startTime,
			EndTime:     endTime,
			EventType:   []string{"single", "recurring"}[rng.IntN(2)],
			TicketType:  ticketType,
			CreatedById: userIds[rng.IntN(len(userIds))],
			CampusId:    campusIds[rng.IntN(len(campusIds))],
			Location:    fake.Address().City(),
			BannerUrl:   fmt.Sprintf("https://picsum.photos/seed/%d/800/400", rng.IntN(200)),
			Categories:  categoriesToAdd,
			TicketTypes: tickets,
		}

		if err := db.Create(&event).Error; err != nil {
			zap.L().Sugar().Errorf("Failed to create event: %v", err)
			return err
		}
	}

	return nil
}

// Utility Functions

func ptr[T any](v T) *T {
	return &v
}

func getUniqueRandomCategories(categories []model.Category, count int, rng *rand.Rand) []model.Category {
	selected := make([]model.Category, 0, count)
	available := make([]model.Category, len(categories))
	copy(available, categories)

	for i := 0; i < count && len(available) > 0; i++ {
		index := rng.IntN(len(available))
		selected = append(selected, available[index])
		available = append(available[:index], available[index+1:]...)
	}

	return selected
}
