package repositories

type SegmentRepository struct {
	baseRepository
}

func NewSegmentRepository(baseRepository baseRepository) *SegmentRepository {
	return &SegmentRepository{baseRepository: baseRepository}
}
