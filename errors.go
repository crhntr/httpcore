package httpcore

import "net/http"

func HandleError(w http.ResponseWriter, errs ...error) []error {
	retErrs := []error{}
	highestStatus := http.StatusInternalServerError
	for _, err := range errs {
		retErrs = append(retErrs, err)
		statusErr, ok := err.(StatusGetter)
		if ok {
			status := statusErr.Status()
			if status > highestStatus {
				highestStatus = status
			}
		}
	}
	if len(retErrs) == 0 {
		return retErrs
	}
	WriteErrorsJSON(w, highestStatus, retErrs...)
	return retErrs
}
