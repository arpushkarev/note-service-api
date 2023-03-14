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
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestGet(t *testing.T) {

	var (
		ctx      = context.Background()
		mockctrl = gomock.NewController(t)

		id        = gofakeit.Int64()
		title     = gofakeit.Word()
		text      = gofakeit.SentenceSimple()
		author    = gofakeit.Name()
		createdAt = gofakeit.Date()
		updatedAt = sql.NullTime{
			Time:  gofakeit.Date(),
			Valid: true,
		}

		req = &desc.GetRequest{
			Id: id,
		}

		validRes = &desc.GetResponse{
			Note: &desc.Note{
				Id: id,
				Note: &desc.NoteInfo{
					Title:  title,
					Text:   text,
					Author: author,
				},
				CreatedAt: timestamppb.New(createdAt),
				UpdatedAt: timestamppb.New(updatedAt.Time),
			},
		}

		repoRes = &model.Note{
			ID: id,
			Info: &model.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		}

		repoErrText = gofakeit.SentenceSimple()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockctrl)

	gomock.InOrder(
		noteMock.EXPECT().Get(ctx, id).Return(repoRes, nil),
		noteMock.EXPECT().Get(ctx, id).Return(nil, repoErr),
	)

	api := note_v1.NewMockNoteV1(note_v1.Implementation{
		NoteService: note.NewMockNoteService(noteMock),
	})

	t.Run("note get success case", func(t *testing.T) {
		res, err := api.Get(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)

	})

	t.Run("note get negative case repo err", func(t *testing.T) {
		_, err := api.Get(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())

	})
}
