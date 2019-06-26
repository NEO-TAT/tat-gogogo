package consts

const (
	// Base url of ntut
	Base      = "https://app.ntut.edu.tw/"
	IndexPage = Base + "index.do"
	Login     = Base + "login.do"
	MainPage  = Base + "myPortal.do"
	// 應用系統頁面
	AptreeListPage = Base + "aptreeList.do"
	// 教務系統頁面
	AptreeAAListPage = AptreeListPage + "?apDn=ou=aa,ou=aproot,o=ldaproot"
	// SSO 登入課程系統
	SsoLoginCourseSystem = Base + "ssoIndex.do?apOu=aa_0010-&apUrl=https://aps.ntut.edu.tw/course/tw/courseSID.jsp"
)
