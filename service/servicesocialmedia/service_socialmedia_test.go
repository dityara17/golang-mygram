package servicesocialmedia

import (
	"errors"
	"testing"
	"time"

	"github.com/arfan21/golang-mygram/entity"
	"github.com/arfan21/golang-mygram/model/modelsocialmedia"
	"github.com/arfan21/golang-mygram/repository/repositoryphoto"
	"github.com/arfan21/golang-mygram/repository/repositorysocialmedia"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

const ENV_TEST_PATH = "../../.env.test"

type ServiceUserTestSuite struct {
	suite.Suite
	repo           *repositorysocialmedia.RepositorySocialMediaMock
	repoPhoto      *repositoryphoto.RepositoryPhotoMock
	srv            ServiceSocialMedia
	defaultPayload modelsocialmedia.Request
}

func TestServiceUser(t *testing.T) {
	err := godotenv.Load(ENV_TEST_PATH)
	assert.NoError(t, err)
	repo := repositorysocialmedia.RepositorySocialMediaMock{}
	repoPhoto := repositoryphoto.RepositoryPhotoMock{}

	defaultPayload := modelsocialmedia.Request{
		Name:           "Test",
		SocialMediaUrl: "https://www.instagram.com/test",
	}

	srv := New(&repo, &repoPhoto)

	testSuite := &ServiceUserTestSuite{
		repo:           &repo,
		repoPhoto:      &repoPhoto,
		srv:            srv,
		defaultPayload: defaultPayload,
	}

	suite.Run(t, testSuite)
}

func (suite *ServiceUserTestSuite) Test_A_CreateSocialMedia() {
	suite.T().Run("Test Create Social Media Success", func(t *testing.T) {
		repoReturn := entity.SocialMedia{
			ID:             1,
			Name:           "Test",
			SocialMediaUrl: "https://www.instagram.com/test",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
			UserID:         1,
		}

		suite.repo.On("Create", mock.Anything).Return(repoReturn, nil)

		createdSocialMedia, err := suite.srv.Create(suite.defaultPayload)
		assert.NoError(t, err)
		assert.NotNil(t, createdSocialMedia)
		assert.Equal(t, createdSocialMedia.ID, repoReturn.ID)
		assert.Equal(t, createdSocialMedia.Name, repoReturn.Name)
		assert.Equal(t, createdSocialMedia.SocialMediaUrl, repoReturn.SocialMediaUrl)
		assert.Equal(t, createdSocialMedia.UserID, repoReturn.UserID)
	})

	suite.T().Run("Test Create Social Media Failed Validation", func(t *testing.T) {
		suite.repo.On("Create", mock.Anything).Return(entity.SocialMedia{}, errors.New("erro"))

		_, err := suite.srv.Create(modelsocialmedia.Request{})
		assert.Error(t, err)
		_, ok := err.(validation.Errors)
		assert.True(t, ok)
	})
}

func (suite *ServiceUserTestSuite) Test_B_GetList() {
	suite.T().Run("Test Get List Success", func(t *testing.T) {
		repoReturn := []entity.SocialMedia{
			{
				ID:             1,
				Name:           "Test",
				SocialMediaUrl: "https://www.instagram.com/test",
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
				UserID:         1,
				User: entity.User{
					ID:       1,
					Username: "test",
				},
			},
		}

		repoPhotoReturn := entity.Photo{
			ID:        1,
			Title:     "Test",
			Caption:   "Test",
			PhotoURL:  "https://www.photo.com/test.jpg",
			UserID:    1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		suite.repo.On("GetList", mock.Anything).Return(repoReturn, nil)
		suite.repoPhoto.On("GetPhotoByUserID", uint(1)).Return(repoPhotoReturn, nil)

		listSocialMedia, err := suite.srv.GetList()
		assert.NoError(t, err)
		assert.NotNil(t, listSocialMedia)
		assert.GreaterOrEqual(t, len(listSocialMedia.SocialMedias), 1)
		assert.Equal(t, listSocialMedia.SocialMedias[0].ID, repoReturn[0].ID)
		assert.Equal(t, listSocialMedia.SocialMedias[0].Name, repoReturn[0].Name)
		assert.Equal(t, listSocialMedia.SocialMedias[0].SocialMediaUrl, repoReturn[0].SocialMediaUrl)
		assert.Equal(t, listSocialMedia.SocialMedias[0].UserID, repoReturn[0].UserID)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.ProfileImageUrl, repoPhotoReturn.PhotoURL)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.Username, repoReturn[0].User.Username)
		assert.Equal(t, listSocialMedia.SocialMedias[0].User.ID, repoReturn[0].User.ID)

	})
}
