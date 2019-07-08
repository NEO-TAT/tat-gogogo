package consts

const (
	// Base url of ntut
	Base = "https://app.ntut.edu.tw/"
	// IndexPage is the idnex url of ntut
	IndexPage = Base + "index.do"
	// Login is login url of ntut
	Login = Base + "login.do"
	// MainPage is the main page of ntut
	MainPage = Base + "myPortal.do"
	// AptreeListPage is application page of ntut
	AptreeListPage = Base + "aptreeList.do"
	// AptreeAAListPage is 教務系統頁面
	AptreeAAListPage = AptreeListPage + "?apDn=ou=aa,ou=aproot,o=ldaproot"
	// SsoLoginCourseSystem is the SSO 登入課程系統
	SsoLoginCourseSystem = Base + "ssoIndex.do?apOu=aa_0010-&apUrl=https://aps.ntut.edu.tw/course/tw/courseSID.jsp"
)
