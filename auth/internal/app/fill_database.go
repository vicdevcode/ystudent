package app

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"os"

	"github.com/vicdevcode/ystudent/auth/internal/dto"
	"github.com/vicdevcode/ystudent/auth/internal/usecase"
)

func fillDatabase(usecases usecase.UseCases) error {
	ctx := context.Background()

	check, err := usecases.FacultyUseCase.FindAll(ctx)
	if err != nil {
		return err
	}
	if len(check) != 0 {
		return nil
	}

	facultyIMI, err := usecases.FacultyUseCase.Create(ctx, dto.CreateFaculty{
		Name: "ИМИ",
	})
	if err != nil {
		return err
	}

	groupIvt202, err := usecases.GroupUseCase.Create(ctx, dto.CreateGroup{
		FacultyID: facultyIMI.ID,
		Name:      "ИВТ-20-2",
	})
	if err != nil {
		return err
	}

	groupIvt201, err := usecases.GroupUseCase.Create(ctx, dto.CreateGroup{
		FacultyID: facultyIMI.ID,
		Name:      "ИВТ-20-1",
	})
	if err != nil {
		return err
	}

	hashedPassword, err := usecases.HashUseCase.HashPassword("123123123")
	if err != nil {
		return err
	}

	for i := 0; i < 20; i++ {
		user, err := usecases.UserUseCase.Create(ctx, dto.CreateUser{
			Fio:      generateFio(),
			Email:    fmt.Sprintf("s%d@gmail.com", i),
			Password: hashedPassword,
		})
		if err != nil {
			return err
		}

		_, err = usecases.StudentUseCase.Create(ctx, dto.CreateStudent{
			UserID:  user.ID,
			Leader:  false,
			GroupID: groupIvt202.ID,
		})
		if err != nil {
			return err
		}
	}

	for i := 20; i < 40; i++ {
		user, err := usecases.UserUseCase.Create(ctx, dto.CreateUser{
			Fio:      generateFio(),
			Email:    fmt.Sprintf("s%d@gmail.com", i),
			Password: hashedPassword,
		})
		if err != nil {
			return err
		}

		_, err = usecases.StudentUseCase.Create(ctx, dto.CreateStudent{
			UserID:  user.ID,
			Leader:  false,
			GroupID: groupIvt201.ID,
		})
		if err != nil {
			return err
		}
	}

	user, err := usecases.UserUseCase.Create(ctx, dto.CreateUser{
		Fio:      generateFio(),
		Email:    "t1@gmail.com",
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	teacher1, err := usecases.TeacherUseCase.Create(ctx, dto.CreateTeacher{
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	_, err = usecases.GroupUseCase.UpdateCurator(ctx, dto.UpdateGroupCurator{
		ID:        groupIvt202.ID,
		CuratorID: teacher1.ID,
	})
	if err != nil {
		return err
	}

	user, err = usecases.UserUseCase.Create(ctx, dto.CreateUser{
		Fio:      generateFio(),
		Email:    "t2@gmail.com",
		Password: hashedPassword,
	})
	if err != nil {
		return err
	}

	teacher2, err := usecases.TeacherUseCase.Create(ctx, dto.CreateTeacher{
		UserID: user.ID,
	})
	if err != nil {
		return err
	}

	_, err = usecases.GroupUseCase.UpdateCurator(ctx, dto.UpdateGroupCurator{
		ID:        groupIvt201.ID,
		CuratorID: teacher2.ID,
	})
	if err != nil {
		return err
	}

	return nil
}

type fio struct {
	MaleSurnames      []string `json:"male_surnames"`
	FemaleSurnames    []string `json:"female_surnames"`
	MaleFirstnames    []string `json:"male_firstnames"`
	FemaleFirstnames  []string `json:"female_firstnames"`
	MaleMiddlenames   []string `json:"male_middlenames"`
	FemaleMiddlenames []string `json:"female_middlenames"`
}

func generateFio() dto.Fio {
	fios, err := os.Open("fio.json")
	if err != nil {
		return dto.Fio{}
	}
	defer fios.Close()

	byteValue, _ := io.ReadAll(fios)
	var result fio
	json.Unmarshal([]byte(byteValue), &result)

	if rand.Intn(2) == 1 {
		return dto.Fio{
			Firstname:  result.MaleFirstnames[rand.Intn(len(result.MaleFirstnames))],
			Surname:    result.MaleSurnames[rand.Intn(len(result.MaleSurnames))],
			Middlename: result.MaleMiddlenames[rand.Intn(len(result.MaleMiddlenames))],
		}
	}
	return dto.Fio{
		Firstname:  result.FemaleFirstnames[rand.Intn(len(result.FemaleFirstnames))],
		Surname:    result.FemaleSurnames[rand.Intn(len(result.FemaleSurnames))],
		Middlename: result.FemaleMiddlenames[rand.Intn(len(result.FemaleMiddlenames))],
	}
}
