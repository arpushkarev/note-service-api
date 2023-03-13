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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func TestUpdate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockctrl = gomock.NewController(t)

		id     = gofakeit.Int64()
		title  = gofakeit.Word()
		text   = gofakeit.SentenceSimple()
		author = gofakeit.Name()

		req = &desc.UpdateRequest{
			Id: id,
			Note: &desc.UpdateNoteInfo{
				Title:  &wrapperspb.StringValue{Value: title},
				Text:   &wrapperspb.StringValue{Value: text},
				Author: &wrapperspb.StringValue{Value: author},
			},
		}

		validRes = &emptypb.Empty{}

		repoReq = &model.UpdateNoteInfo{
			Title: sql.NullString{
				String: title,
				Valid:  true,
			},
			Text: sql.NullString{
				String: text,
				Valid:  true,
			},
			Author: sql.NullString{
				String: author,
				Valid:  true,
			},
		}

		repoErrText = gofakeit.SentenceSimple()
		repoErr     = errors.New(repoErrText)
	)

	noteMock := noteMocks.NewMockRepository(mockctrl)

	gomock.InOrder(
		noteMock.EXPECT().Update(ctx, id, repoReq).Return(nil),
		noteMock.EXPECT().Update(ctx, id, repoReq).Return(repoErr),
	)

	api := note_v1.NewMockNoteV1(note_v1.Implementation{
		NoteService: note.NewMockNoteService(noteMock),
	})

	t.Run("note update success case", func(t *testing.T) {
		res, err := api.Update(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validRes, res)
	})

	t.Run("note update negative case repo err", func(t *testing.T) {
		_, err := api.Update(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
