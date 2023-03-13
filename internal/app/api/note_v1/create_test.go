package note_v1_test

import (
	"context"
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
)

func TestCreate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id     = gofakeit.Int64()
		title  = gofakeit.Word()
		text   = gofakeit.SentenceSimple()
		author = gofakeit.Name()

		req = &desc.CreateRequest{
			Note: &desc.NoteInfo{
				Title:  title,
				Text:   text,
				Author: author,
			},
		}

		validRes = &desc.CreateResponse{
			Id: id,
		}

		repoReq = &model.NoteInfo{
			Title:  title,
			Text:   text,
			Author: author,
		}

		repoErrText = gofakeit.SentenceSimple()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockCtrl)

	gomock.InOrder(
		noteMock.EXPECT().Create(ctx, repoReq).Return(id, nil),
		noteMock.EXPECT().Create(ctx, repoReq).Return(int64(0), repoErr),
	)

	api := note_v1.NewMockNoteV1(note_v1.Implementation{
		NoteService: note.NewMockNoteService(noteMock),
	})

	t.Run(" note create success case", func(t *testing.T) {
		res, err := api.Create(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note create negative case repo err", func(t *testing.T) {
		_, err := api.Create(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
