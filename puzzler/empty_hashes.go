package puzzler



type emptyHashStruct struct {
	emptymd5     string
	emptysha1    string
	emptysha256  string
	emptysha512  string
	emptysha3256 string
}

func (hashList *emptyHashStruct) fillEmptyHashStruct() {

	hashList.emptymd5 = Md5Hash("")
	hashList.emptysha1 = Sha1Hash("")
	hashList.emptysha256 = Sha256Hash("")
	hashList.emptysha512 = Sha512Hash("")
	hashList.emptysha3256 = Sha3256("")

}
