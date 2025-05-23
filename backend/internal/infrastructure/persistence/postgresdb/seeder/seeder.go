package postgresdb

import (
	"fmt"
	"os"

	"github.com/HappyNaCl/Cavent/backend/internal/domain/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
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

	campusLogo := fmt.Sprintf("%s/storage/v1/object/public/%s/assets/binus.png", os.Getenv("SUPABASE_PROJECT_URL"), os.Getenv("SUPABASE_BUCKET_NAME"))
	campus := &model.Campus{
		Id: uuid.New(),
		Name: "Binus University",
		LogoUrl: campusLogo,
		Description: "Best University of The West",
		InviteCode: "AAAAAA",
	}

	err := db.Create(campus).Error
	if err != nil {
		zap.L().Sugar().Fatal("Failed to seed campus!")
		return err
	}

	for _, item := range tagData {
		if err := seedTagTypeWithTags(db, item.TypeName, item.Tags); err != nil {
			zap.L().Sugar().Fatalf("Failed to see tag %v", item)
			return err
		}
	}


	zap.L().Sugar().Info("Database seeded successfully!")
	return nil
}

func seedTagTypeWithTags(db *gorm.DB, typeName string, tagNames []string) error {
	tagType := model.CategoryType{
		Id:   uuid.New(),
		Name: typeName,
	}

	if err := db.Create(&tagType).Error; err != nil {
		return err
	}

	var tags []model.Category
	for _, name := range tagNames {
		tags = append(tags, model.Category{
			Id:        uuid.New(),
			Name:      name,
			CategoryTypeId: tagType.Id,
		})
	}

	return insertTags(db, tags)
}

func insertTags(db *gorm.DB, tags []model.Category) error {
	for _, tag := range tags {
		if err := db.Create(&tag).Error; err != nil {
			return err
		}
	}
	return nil
}
