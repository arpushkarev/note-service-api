package note_v1_test

import (
	"context"
	"database/sql"
	"errors"
	"testing"

	"github.com/arpushkarev/note-service-api/internal/app/api/note_v1"
	"github.com/arpushkarev/note-service-api/internal/model"
	noteMocks "github.com/arpushkarev/note-service-api/internal/repository/note/mocks"
	"github.com/arpushkarev/note-service-api/internal/service/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGetAll(t *testing.T) {
	var (
		ctx      = context.Background()
		mockctrl = gomock.NewController(t)

		req = &emptypb.Empty{}

		validResList []*desc.Note

		repoRes []*model.Note

		repoErrText = gofakeit.SentenceSimple()
		repoErr     = errors.New(repoErrText)
	)

	for i := 0; i < 3; i++ {
		id := gofakeit.Int64()
		title := gofakeit.Word()
		text := gofakeit.SentenceSimple()
		author := gofakeit.Name()
		createdAt := gofakeit.Date()
		updatedAt := sql.NullTime{
			Time:  gofakeit.Date(),
			Valid: true,
		}

		repoRes = append(repoRes, &model.Note{
			ID: id,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})

		validResList = append(validResList, &desc.Note{
			Id: id,
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: timestamppb.New(createdAt),
			UpdatedAt: timestamppb.New(updatedAt.Time),
		})
	}

	validRes := &desc.GetAllResponse{Notes: validResList}

	noteMock := noteMocks.NewMockRepository(mockctrl)

	gomock.InOrder(
		noteMock.EXPECT().GetAll(ctx).Return(repoRes, nil),
		noteMock.EXPECT().GetAll(ctx).Return(nil, repoErr),
	)

	api := note_v1.NewMockNoteV1(note_v1.Implementation{
		NoteService: note.NewMockNoteService(noteMock),
	})

	t.Run("note get all success case", func(t *testing.T) {
		res, err := api.GetAll(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note get all negative case repo err", func(t *testing.T) {
		_, err := api.GetAll(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
