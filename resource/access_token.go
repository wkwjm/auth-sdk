package resource

type AccessToken struct {
	Raw       string
	Header    map[string]interface{}
	Sub       string
	Scopes    []string
	GrantType string
	ClientId  string
}

func (token *AccessToken) HasScopes(scopes ...string) bool {
	return contains(token.Scopes, scopes...)
}

func (token *AccessToken) HasGrantType(grantTypes ...string) bool {
	return contains(grantTypes, token.GrantType)
}

func contains(s1 []string, s2 ...string) bool {
	// 创建一个map来跟踪s1中的元素
	seen := make(map[string]bool)
	for _, str := range s1 {
		seen[str] = true
	}

	// 检查s2中的每个元素是否在s1中
	for _, str := range s2 {
		if !seen[str] {
			// 如果s2中的元素在s1中找不到，则返回false
			return false
		}
	}

	// 如果s2中的所有元素都在s1中找到了，则返回true
	return true
}
