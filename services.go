package main

type DatamartService struct {
	repo DatamartRepository
}

func NewDatamartService(repo DatamartRepository) DatamartService {
	return DatamartService{repo}
}

func (s DatamartService) GenerateDatamart1() error {
	data, _ := s.repo.GetDatamart1()

	s.repo.InserDatamart1(data)

	return nil
}
