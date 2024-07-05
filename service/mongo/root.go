package mongo

import (
	"errors"
	pM "go.mongodb.org/mongo-driver/mongo"
	"log"
	"plabfootball/repository"
	"plabfootball/types"
	"plabfootball/types/schema"
)

type MService struct {
	repository *repository.Repository
}

func NewMService(repository *repository.Repository) *MService {
	r := &MService{
		repository: repository,
	}
	return r
}

func (m *MService) View(sch string, region, sex int) (*schema.Stadium, error) {
	if res, err := m.repository.Mongo.View(sch, region, sex); err != nil {
		log.Println("Failed To Call View Data", "err :", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (m *MService) ViewAll() ([]*schema.Stadium, error) {
	if res, err := m.repository.Mongo.ViewAll(); err != nil {
		log.Println("Failed To Call ViewAll Data", "err :", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (m *MService) Upsert(sch string, sex, region int, upsert types.AddReq) (*schema.Stadium, error) {
	if res, err := m.repository.Mongo.Upsert(sch, sex, region, upsert); err != nil {
		log.Println("Failed To Call Upsert Data", "err :", err)
		return nil, err
	} else {
		return res, nil
	}
}

func (m *MService) Delete(sch string, sex, region int) error {
	if err := m.repository.Mongo.Delete(sch, region, sex); err != nil {
		log.Println("Failed To Call Delete Data", "err :", err)
		return err
	} else {
		return nil
	}
}

func (m *MService) Add(sch string, sex, region int) error {
	// 데이터가 있는지 먼저 검색.
	if _, err := m.repository.Mongo.View(sch, region, sex); err != nil {
		if errors.Is(err, pM.ErrNoDocuments) {
			// document가 없으면 데이터를 추가합니다.
			if err = m.repository.Mongo.Add(sch, sex, region); err != nil {
				log.Println("Failed To Call Add Data", "err :", err)
				return err
			} else {
				return nil
			}
		}
	}
	return errors.New("Document가 이미 존재합니다.")
}

func (m *MService) GetGirlUser(sch string, region, sex int) (*types.GirlUrlRes, error) {
	if res, err := m.repository.Mongo.View(sch, region, sex); err != nil {
		log.Println("Failed To Call View Data", "err :", err)
		return nil, err
	} else if urls, err := getGirlStadium(res.URL); err != nil {
		return urls, err
	} else {
		return urls, nil
	}
}
