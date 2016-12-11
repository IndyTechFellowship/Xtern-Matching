package services

import "golang.org/x/net/context"

func ExportStudents(ctx context.Context) (string, error) {
	students, err := GetStudents(ctx)
	if err != nil {
		return "", err
	}
	//Won't Marshal because of array.  Will convert array in model to string.
	output, err := Marshal(students)
	if err != nil {
		return "", err
	}
	return string(output), nil
}
