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

// NewMService 는 Mongo repository와 MongoRouter를 연결하는 다리인 MService 객체를 생성합니다.
func NewMService(repository *repository.Repository) *MService {
	r := &MService{
		repository: repository,
	}
	return r
}

// View 는 MongoRouter에서 보낸 데이터를 repository에 View에 전달합니다.
func (m *MService) View(sch string, region, sex int) (*schema.Stadium, error) {
	if res, err := m.repository.Mongo.View(sch, region, sex); err != nil {
		log.Println("Failed To Call view Data", "err :", err)
		return nil, err
	} else {
		return res, nil
	}
}

// ViewAll 는 MongoRouter에서 보낸 데이터를 repository의 ViewAll에 전달합니다.
func (m *MService) ViewAll() ([]*schema.Stadium, error) {
	if res, err := m.repository.Mongo.ViewAll(); err != nil {
		log.Println("Failed To Call ViewAll Data", "err :", err)
		return nil, err
	} else {
		return res, nil
	}
}

// Upsert 는 MongoRouter에서 보낸 데이터를 repository의 Upsert에 전달합니다.
func (m *MService) Upsert(sch string, sex, region int, upsert types.AddReq) (*schema.Stadium, error) {
	if _, err := m.repository.Mongo.View(upsert.Sch, upsert.Region, upsert.Sex); err == nil {
		log.Println("Data that you want to insert is already exist", "err :", err)
		return nil, errors.New("Data that you want to insert is already exist")
	} else if res, err := m.repository.Mongo.Upsert(sch, sex, region, upsert); err != nil {
		log.Println("Failed To Call Update Data, but created new document", "err :", err)
		return nil, errors.New("Created new document")
	} else {
		return res, nil
	}
}

// Delete 는 MongoRouter에서 보낸 데이터를 repository의 Delete에 전달합니다.
func (m *MService) Delete(sch string, sex, region int) error {
	if err := m.repository.Mongo.Delete(sch, region, sex); err != nil {
		log.Println("Failed To Call Delete Data", "err :", err)
		return err
	} else {
		return nil
	}
}

// Add 는 MongoRouter에서 보낸 데이터를 repository의 View에 전달하여 데이터 중복여부를 체크후 Add에 전달합니다.
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

// GetGirlUser 는 MongoRouter에서 보낸 데이터를 repository의 View에 전달 데이터 여부를 확인 후 getGirlStadium 함수를 호출하여 경기장을 가져옵니다.
func (m *MService) GetGirlUser(sch string, region, sex int) (*types.GirlUrlRes, error) {
	if res, err := m.repository.Mongo.View(sch, region, sex); err != nil {
		log.Println("Failed To Call view Data", "err :", err)
		return nil, err
	} else if urls, err := getGirlStadium(res.URL); err != nil {
		return urls, err
	} else {
		return urls, nil
	}
}
