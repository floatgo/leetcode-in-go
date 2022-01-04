// APPROACH 1
func stringSorter(s string) string {
    s2 := strings.Split(s, "")
    sort.Strings(s2)
    return strings.Join(s2, "")
}

func isAnagram(s string, t string) bool {
    ss := stringSorter(s)
    st := stringSorter(t)
    
    return ss == st
}


// APPROACH 2
func isAnagram(s string, t string) bool {
    
    if len(s) != len(t) {
        return false
    }
    
    m_s := map[byte]int{}
    m_t := map[byte]int{}
    
    for i := 0; i < len(s); i++ {
        if _, ok := m_s[s[i]]; ok {
            m_s[s[i]] += 1
        } else {
            m_s[s[i]] = 0
        }
    }
    
    for i := 0; i < len(t); i++ {
        if _, ok := m_t[t[i]]; ok {
            m_t[t[i]] += 1
        } else {
            m_t[t[i]] = 0
        }
    }
    
    for k, _ := range m_s {
        if _, ok := m_t[k]; !ok {
            return false
        }
        
        if m_s[k] != m_t[k] {
            return false
        }
    }
    return true
}

