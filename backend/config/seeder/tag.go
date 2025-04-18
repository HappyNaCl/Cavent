package seeder

import (
	"fmt"

	"github.com/HappyNaCl/Cavent/backend/domain/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagSeeder struct{}

func NewTagSeeder() *TagSeeder {
	return &TagSeeder{}
}

func (s *TagSeeder) Seed(db *gorm.DB) error {
	seedData := []struct {
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

	for _, item := range seedData {
		if err := seedTagTypeWithTags(db, item.TypeName, item.Tags); err != nil {
			return fmt.Errorf("failed to seed tag type %s: %w", item.TypeName, err)
		}
	}

	return nil
}

func seedTagTypeWithTags(db *gorm.DB, typeName string, tagNames []string) error {
	tagType := model.TagType{
		Id:   uuid.NewString(),
		Name: typeName,
	}

	if err := db.Create(&tagType).Error; err != nil {
		return err
	}

	var tags []model.Tag
	for _, name := range tagNames {
		tags = append(tags, model.Tag{
			Id:        uuid.NewString(),
			Name:      name,
			TagTypeId: tagType.Id,
		})
	}

	return insertTags(db, tags)
}

func insertTags(db *gorm.DB, tags []model.Tag) error {
	for _, tag := range tags {
		if err := db.Create(&tag).Error; err != nil {
			return err
		}
	}
	return nil
}
