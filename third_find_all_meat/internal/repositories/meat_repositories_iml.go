package repositories

import (
	"backend/third_find_all_meat/internal/core/port"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
)

type meatRepository struct {
}

func NewMeatRepository() port.IMeatRepository {
	return &meatRepository{}
}

func (r *meatRepository) Get(ctx context.Context, index string) (string, error) {
	requestURL := fmt.Sprintf("https://baconipsum.com/api/?type=meat-and-filler&paras=%v&format=text", index)
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		os.Exit(1)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return "", err
	}
	defer res.Body.Close()

	fmt.Printf("client: got response!\n")
	fmt.Printf("client: status code: %d\n", res.StatusCode)

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("status code: %d", res.StatusCode)
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return "", err
	}
	return string(resBody), nil
}
