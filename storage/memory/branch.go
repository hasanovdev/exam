package memory

import (
	"encoding/json"
	"errors"
	"log"
	"os"

	"github.com/google/uuid"

	"exam/models"
)

type BranchRepo struct {
	fileName string
	file     *os.File
}

func NewBranchRepo(fileName string, file *os.File) *BranchRepo {
	return &BranchRepo{
		fileName: fileName,
		file:     file,
	}
}

func (u *BranchRepo) Create(req *models.CreateBranch) (*models.Branch, error) {

	Branches, err := u.read()
	if err != nil {
		return nil, err
	}

	var (
		id     = uuid.New().String()
		Branch = models.Branch{
			Id:   id,
			Name: req.Name,
		}
	)
	Branches[id] = Branch

	err = u.write(Branches)
	if err != nil {
		return nil, err
	}

	return &Branch, nil
}

func (u *BranchRepo) GetById(req *models.BranchPrimaryKey) (*models.Branch, error) {

	Branches, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Branches[req.Id]; !ok {
		return nil, errors.New("Branch not found")
	}
	Branch := Branches[req.Id]

	return &Branch, nil
}

func (u *BranchRepo) GetList(req *models.BranchGetListRequest) (*models.BranchGetListResponse, error) {

	var resp = &models.BranchGetListResponse{}
	resp.Branches = []*models.Branch{}

	BranchMap, err := u.read()
	if err != nil {
		return nil, err
	}

	resp.Count = len(BranchMap)
	for _, val := range BranchMap {
		Branch := val
		resp.Branches = append(resp.Branches, &Branch)
	}

	return resp, nil
}

func (u *BranchRepo) Update(req *models.UpdateBranch) (*models.Branch, error) {

	Branches, err := u.read()
	if err != nil {
		return nil, err
	}

	if _, ok := Branches[req.Id]; !ok {
		return nil, errors.New("Category not found")
	}

	Branches[req.Id] = models.Branch{
		Id:   req.Id,
		Name: req.Name,
	}

	err = u.write(Branches)
	if err != nil {
		return nil, err
	}
	Category := Branches[req.Id]

	return &Category, nil
}

func (u *BranchRepo) Delete(req *models.BranchPrimaryKey) error {

	Branches, err := u.read()
	if err != nil {
		return err
	}

	delete(Branches, req.Id)

	err = u.write(Branches)
	if err != nil {
		return err
	}

	return nil
}

func (u *BranchRepo) read() (map[string]models.Branch, error) {
	var (
		Branches  []*models.Branch
		BranchMap = make(map[string]models.Branch)
	)

	data, err := os.ReadFile(u.fileName)
	if err != nil {
		log.Printf("Error while Read data: %+v", err)
		return nil, err
	}

	err = json.Unmarshal(data, &Branches)
	if err != nil {
		log.Printf("Error while Unmarshal data: %+v", err)
		return nil, err
	}

	for _, branch := range Branches {
		BranchMap[branch.Id] = *branch
	}

	return BranchMap, nil
}

func (u *BranchRepo) write(BranchMap map[string]models.Branch) error {

	var Branches []models.Branch

	for _, val := range BranchMap {
		Branches = append(Branches, val)
	}

	body, err := json.MarshalIndent(Branches, "", "	")
	if err != nil {
		return err
	}

	err = os.WriteFile(u.fileName, body, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}
