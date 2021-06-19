//
package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMandiriWorker_Run(t *testing.T) {
	t.Skip()
	w := &mandiriWorker{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)
	defer cancel()
	rates, err := w.Run(ctx)
	b, _ := json.Marshal(rates)
	require.Nil(t, err)
	require.Greater(t, len(rates), 0)
	fmt.Println("marshaled", string(b))
}
