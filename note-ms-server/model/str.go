package model

import (
  "regexp"
  "strings"
)

func TrimWholeString(s string) string {
  s = strings.TrimSpace(s)
  re := regexp.MustCompile(`\s`)
  s = re.ReplaceAllString(s, "")

  return s
}

func TrimMutipleSpaceWholeString(s string) string {
  s = strings.TrimSpace(s)
  re := regexp.MustCompile(`\s+`)
  s = re.ReplaceAllString(s, " ")

  return s
}

func TrimFullText(s string) string {
  temp := strings.Split(s, "\n")
  var result []string
  for _, v := range temp {
    trimStr := TrimMutipleSpaceWholeString(v)
    if len(trimStr) > 0 {
      result = append(result, trimStr)
    }
  }

  return strings.Join(result[:], "\n")
}

func SubString(s string, l int) string {
  rs := []rune(s)

  if l > len(rs) {
    l = len(rs)
  }

  return string(rs[:l])
}

func NormSubString(s string, l int) string {
  rs := []rune(s)

  i := 0
  j := -1
  curLen := 0
  wordLen := 0
  for i < len(rs) {
    c := runeWidth(rs[i])
    curLen = curLen + c
    if isStopWord(rs[i]) {
      if i-1 > -1 && rs[i-1] != rune('.') {
        j = i
      }
      wordLen = 0
    } else {
      wordLen = wordLen + c
    }
    // 返回
    if curLen > l {
      // 没有stopWord的情况
      if j == -1 {
        if i == len(rs)-1 {
          return string(rs[:i+1])
        } else {
          return string(rs[:i+1]) + "..."
        }
      }
      // 有stopWord的情况并且wordLen<8
      if j != -1 && wordLen <= 8 {
        return string(rs[:j]) + "..."
      }
      // 有stopWord的情况并且wordLen>8
      if j != -1 && wordLen > 8 {
        return string(rs[:i+1]) + "..."
      }
    }
    i = i + 1
  }
  return s
}

func NormSubDomain(s string, l int) string {
  rs := []rune(s)

  i := 0
  j := -1
  curLen := 0
  wordLen := 0
  for i < len(rs) {
    c := runeWidth(rs[i])
    curLen = curLen + c
    if isStopWord(rs[i]) {
      if i-1 > -1 && rs[i-1] != rune('.') {
        j = i
      }
      wordLen = 0
    } else {
      wordLen = wordLen + c
    }
    // 返回
    if curLen > l {
      // 没有stopWord的情况
      if j == -1 {
        if i == len(rs)-1 {
          return string(rs[:i+1])
        } else {
          return string(rs[:i+1]) + "..."
        }
      }
      // 有stopWord的情况并且wordLen<8
      if j != -1 && wordLen <= 3 {
        return string(rs[:j]) + "..."
      }
      // 有stopWord的情况并且wordLen>8
      if j != -1 && wordLen > 3 {
        return string(rs[:i+1]) + "..."
      }
    }
    i = i + 1
  }
  return s
}

func runeWidth(c rune) int {
  if (c>>11) > 0 || (c>>7) > 0 {
    return 2
  } else {
    return 1
  }
}

func isStopWord(c rune) bool {
  p := map[rune]bool{rune('['): true, rune('.'): true, rune(','): true, rune('?'): true, rune('('): true, rune(')'): true, rune('!'): true, rune('\''): true, rune('"'): true, rune(':'): true, rune(';'): true, rune('-'): true, rune('—'): true, rune('。'): true, rune('？'): true, rune('！'): true, rune('，'): true, rune('、'): true, rune('；'): true, rune('：'): true, rune('“'): true, rune('”'): true, rune('﹃'): true, rune('﹄'): true, rune('﹁'): true, rune('﹂'): true, rune('（'): true, rune('）'): true, rune('［'): true, rune('］'): true, rune('〔'): true, rune('〕'): true, rune('【'): true, rune('】'): true, rune('…'): true, rune('《'): true, rune('》'): true, rune('〈'): true, rune('〉'): true, rune('﹏'): true, rune('＿'): true, rune(' '): true}
  if _, ok := p[c]; ok {
    return true
  } else {
    return false
  }
}

func normSubStringFullText(s string, l int) string {
  rs := []rune(s)

  i := 0
  j := -1
  curLen := 0
  wordLen := 0
  for i < len(rs) {
    c := 1
    curLen = curLen + c
    if isStopWord(rs[i]) {
      if i-1 > -1 && rs[i-1] != rune('.') {
        j = i
      }
      wordLen = 0
    } else {
      wordLen = wordLen + c
    }
    // 返回
    if curLen > l {
      // 没有stopWord的情况
      if j == -1 {
        if i == len(rs)-1 {
          return string(rs[:i+1])
        } else {
          return string(rs[:i+1]) + "..."
        }
      }
      // 有stopWord的情况并且wordLen<6
      if j != -1 && wordLen <= 6 {
        return string(rs[:j]) + "..."
      }
      // 有stopWord的情况并且wordLen>6
      if j != -1 && wordLen > 6 {
        return string(rs[:i+1]) + "..."
      }
    }
    i = i + 1
  }
  return s
}

func subStringSlice(s []string, l int) []string {
  var result []string

  curLen := 0

  for _, v := range s {
    trimStr := TrimMutipleSpaceWholeString(v)
    rs := []rune(trimStr)
    if curLen+len(rs) > l {
      if l-curLen > 8 {
        result = append(result, normSubStringFullText(trimStr, l-curLen))
      }
      return result
    }
    result = append(result, trimStr)
    curLen = curLen + len(rs)
  }

  return result
}

func NormSubFullText(s string, l int) string {
  temp := strings.Split(s, "\n")

  strSlice := subStringSlice(temp, l)

  return strings.Join(strSlice[:], "\n")
}

func StringLength(s string) int {
  rs := []rune(s)

  return len(rs)
}
