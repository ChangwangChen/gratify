package buffer

type ReferenceCounted interface {
   RefCnt() int
   Retain() ReferenceCounted
   RetainByCount(increment int) ReferenceCounted
   Touch() ReferenceCounted
   TouchByHint(hint interface{})
   Release() ReferenceCounted
   ReleaseByCount(decrement int) ReferenceCounted
}