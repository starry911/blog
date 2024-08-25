package uuidStr

import "github.com/google/uuid"

// UUID1 生成V1版本的UUID
func UUID1() (string, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// UUID2G 生成V2版本的UUID，GID版本
func UUID2G() (string, error) {
	id, err := uuid.NewDCEGroup()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// UUID2P 生成V2版本的UUID，POSIX UID版本
func UUID2P() (string, error) {
	id, err := uuid.NewDCEPerson()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

// UUID3 生成V3版本的UUID
func UUID3(name string) (string, error) {
	id, err := uuid.NewDCEPerson()
	if err != nil {
		return "", err
	}
	id2 := uuid.NewMD5(id, []byte(name))
	return id2.String(), nil
}

// UUID4 生成V5版本的UUID
func UUID4() string {
	id := uuid.New()
	return id.String()
}

// UUID5 生成V5版本的UUID
func UUID5(name string) (string, error) {
	id, err := uuid.NewDCEPerson()
	if err != nil {
		return "", err
	}
	id2 := uuid.NewSHA1(id, []byte(name))
	return id2.String(), nil
}
