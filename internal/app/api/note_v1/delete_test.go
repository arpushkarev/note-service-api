package note_v1_test

import (
	"context"
	"errors"
	"testing"

	"github.com/arpushkarev/note-service-api/internal/app/api/note_v1"
	noteMocks "github.com/arpushkarev/note-service-api/internal/repository/note/mocks"
	"github.com/arpushkarev/note-service-api/internal/service/note"
	desc "github.com/arpushkarev/note-service-api/pkg/note_v1"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestDelete(t *testing.T) {
	var (
		ctx      = context.Background()
		mockctrl = gomock.NewController(t)

		id = gofakeit.Int64()

		req = &desc.DeleteRequest{
			Id: id,
		}

		validRes = &emptypb.Empty{}

		repoErrText = gofakeit.SentenceSimple()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockctrl)

	gomock.InOrder(
		noteMock.EXPECT().Delete(ctx, id).Return(nil),
		noteMock.EXPECT().Delete(ctx, id).Return(repoErr),
	)

	api := note_v1.NewMockNoteV1(note_v1.Implementation{
		NoteService: note.NewMockNoteService(noteMock),
	})

	t.Run("note delete success case", func(t *testing.T) {
		res, err := api.Delete(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note delete negative case repo err", func(t *testing.T) {
		_, err := api.Delete(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
