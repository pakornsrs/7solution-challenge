package authutil

import (
	"context"
	"os"
	timeutil "pakornssn/7solution-challenge/pkg/time-util"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestGenerateAuthenticationToken(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		// arrange
		os.Setenv("JWT_SECRET_KEY", "aJm5+7P3tFbDqZHTq8gW4nqoxzWJ3i28bM6oW8Zz9Xg=")
		os.Setenv("ISSUER", "user-service")
		os.Setenv("AUDIENCE", "user-service-client")

		mockTimeUtc := time.Date(2030, 7, 25, 22, 58, 20, 0, time.UTC)
		timeutil.TimeNowUte = func() time.Time {
			return mockTimeUtc
		}

		userId := "mock-user-id"

		duration := 15 * time.Minute

		ctx := context.Background()

		// act
		resp, err := GenerateAuthenticationToken(ctx, userId, duration)

		// assert
		assert.NoError(t, err)

		expectToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJ1c2VyLXNlcnZpY2UiLCJzdWIiOiJtb2NrLXVzZXItaWQiLCJhdWQiOlsidXNlci1zZXJ2aWNlLWNsaWVudCJdLCJleHAiOjE5MTEyNTE2MDAsIm5iZiI6MTkxMTI1MDcwMCwiaWF0IjoxOTExMjUwNzAwfQ.H7AwP_WdaV-6I1iI-OqxkAhnQPvE2cbt7WqDyB50xWk"
		assert.Equal(t, expectToken, resp)
	})
}
