package services

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/iqbaludinm/library-api/models"
	"github.com/iqbaludinm/library-api/repositories/mocks"
	"github.com/stretchr/testify/assert"
)

func TestBookService_GetBooks(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult []models.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase
	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	testTable = append(testTable, testCase{
		name:      "Get Books Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]models.Book{
				{
					ID:        1,
					NameBook:  "Filosofi Teras",
					Author:    "Henry Manampiring",
					CreatedAt: now,
					UpdatedAt: now,
				},
				{
					ID:        2,
					NameBook:  "Hujan",
					Author:    "Tere Liye",
					CreatedAt: now,
					UpdatedAt: now,
				},
			}, nil).Times(1)
		},
		expectedResult: []models.Book{
			{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				CreatedAt: now,
				UpdatedAt: now,
			},
			{
				ID:        2,
				NameBook:  "Hujan",
				Author:    "Tere Liye",
				CreatedAt: now,
				UpdatedAt: now,
			},
		},
	})

	testTable = append(testTable, testCase{
		name:          "Get Books Failed",
		wantError:     true,
		expectedError: errors.New("data not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBooks().Return([]models.Book{}, errors.New("data not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := BaseService{
				repo: bookRepo,
			}

			res, err := service.GetBooks()

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}

func TestBookService_GetBookById(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult models.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase
	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	testTable = append(testTable, testCase{
		name:      "Get Book By Id Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				CreatedAt: now,
				UpdatedAt: now,
			}, nil).Times(1)
		},
		expectedResult: models.Book{
			ID:        1,
			NameBook:  "Filosofi Teras",
			Author:    "Henry Manampiring",
			CreatedAt: now,
			UpdatedAt: now,
		},
	})

	testTable = append(testTable, testCase{
		name:          "Get Book By Id Failed",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().GetBookById(gomock.Any()).Return(models.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := BaseService{
				repo: bookRepo,
			}

			res, err := service.GetBookById(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}

func TestBookService_CreateBook(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult models.Book
		input          models.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	testTable = append(testTable, testCase{
		name:      "Create Book Success",
		wantError: false,
		input: models.Book{
			NameBook: "Filosofi Teras",
			Author:   "Henry Manampiring",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				CreatedAt: now,
				UpdatedAt: now,
			}, nil).Times(1)
		},
		expectedResult: models.Book{
			ID:        1,
			NameBook:  "Filosofi Teras",
			Author:    "Henry Manampiring",
			CreatedAt: now,
			UpdatedAt: now,
		},
	})

	testTable = append(testTable, testCase{
		name:      "Create Book Failed",
		wantError: true,
		input: models.Book{
			NameBook: "Filosofi Teras",
			Author:   "Henry Manampiring",
		},
		expectedError: errors.New("failed to store in database"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().CreateBook(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				CreatedAt: now,
				UpdatedAt: now,
			}, errors.New("failed to store in database")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := BaseService{
				repo: bookRepo,
			}

			res, err := service.CreateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}

func TestBookService_UpdateBook(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult models.Book
		input          models.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase

	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	testTable = append(testTable, testCase{
		name:      "Update Book Success",
		wantError: false,
		input: models.Book{
			NameBook: "Filosofi Teras",
			Author:   "Henry Manampiring",
		},
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				UpdatedAt: now,
			}, nil).Times(1)
		},
		expectedResult: models.Book{
			ID:        1,
			NameBook:  "Filosofi Teras",
			Author:    "Henry Manampiring",
			UpdatedAt: now,
		},
	})

	testTable = append(testTable, testCase{
		name:      "Update Book Failed",
		wantError: true,
		input: models.Book{
			NameBook: "Filosofi Teras",
			Author:   "Henry Manampiring",
		},
		expectedError: errors.New("failed to update"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().UpdateBook(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				CreatedAt: now,
				UpdatedAt: now,
			}, errors.New("failed to update")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := BaseService{
				repo: bookRepo,
			}

			res, err := service.UpdateBook(testCase.input)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}

func TestBookService_DeleteBook(t *testing.T) {
	type testCase struct {
		name           string
		expectedResult models.Book
		expectedError  error
		wantError      bool
		onBookRepo     func(mock *mocks.MockBookRepo)
	}

	var testTable []testCase
	now := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	testTable = append(testTable, testCase{
		name:      "Delete Book Success",
		wantError: false,
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(models.Book{
				ID:        1,
				NameBook:  "Filosofi Teras",
				Author:    "Henry Manampiring",
				DeletedAt: &now,
			}, nil).Times(1)
		},
		expectedResult: models.Book{
			ID:        1,
			NameBook:  "Filosofi Teras",
			Author:    "Henry Manampiring",
			DeletedAt: &now,
		},
	})

	testTable = append(testTable, testCase{
		name:          "Delete Book Failed",
		wantError:     true,
		expectedError: errors.New("record not found"),
		onBookRepo: func(mock *mocks.MockBookRepo) {
			mock.EXPECT().DeleteBook(gomock.Any()).Return(models.Book{}, errors.New("record not found")).Times(1)
		},
	})

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			mockController := gomock.NewController(t)

			bookRepo := mocks.NewMockBookRepo(mockController)

			if testCase.onBookRepo != nil {
				testCase.onBookRepo(bookRepo)
			}

			service := BaseService{
				repo: bookRepo,
			}

			res, err := service.DeleteBook(1)

			if testCase.wantError {
				assert.EqualError(t, err, testCase.expectedError.Error())
			} else {
				assert.Nil(t, err)
				assert.Equal(t, testCase.expectedResult, res)
			}

		})
	}
}
