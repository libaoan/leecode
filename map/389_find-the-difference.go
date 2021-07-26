package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Printf("%s\n", string(findTheDifference("aa", "aaa")))
	fmt.Printf("%s\n", string(findTheDifference("abcd", "abcde")))

	fmt.Printf("%s\n", string(findTheDifference2("abcd", "abcde")))
	fmt.Printf("%s\n", string(findTheDifference2("aa", "aaa")))

	s := "ymbgaraibkfmvocpizdydugvalagaivdbfsfbepeyccqfepzvtpyxtbadkhmwmoswrcxnargtlswqemafandgkmydtimuzvjwxvlfwlhvkrgcsithaqlcvrihrwqkpjdhgfgreqoxzfvhjzojhghfwbvpfzectwwhexthbsndovxejsntmjihchaotbgcysfdaojkjldprwyrnischrgmtvjcorypvopfmegizfkvudubnejzfqffvgdoxohuinkyygbdzmshvyqyhsozwvlhevfepdvafgkqpkmcsikfyxczcovrmwqxxbnhfzcjjcpgzjjfateajnnvlbwhyppdleahgaypxidkpwmfqwqyofwdqgxhjaxvyrzupfwesmxbjszolgwqvfiozofncbohduqgiswuiyddmwlwubetyaummenkdfptjczxemryuotrrymrfdxtrebpbjtpnuhsbnovhectpjhfhahbqrfbyxggobsweefcwxpqsspyssrmdhuelkkvyjxswjwofngpwfxvknkjviiavorwyfzlnktmfwxkvwkrwdcxjfzikdyswsuxegmhtnxjraqrdchaauazfhtklxsksbhwgjphgbasfnlwqwukprgvihntsyymdrfovaszjywuqygpvjtvlsvvqbvzsmgweiayhlubnbsitvfxawhfmfiatxvqrcwjshvovxknnxnyyfexqycrlyksderlqarqhkxyaqwlwoqcribumrqjtelhwdvaiysgjlvksrfvjlcaiwrirtkkxbwgicyhvakxgdjwnwmubkiazdjkfmotglclqndqjxethoutvjchjbkoasnnfbgrnycucfpeovruguzumgmgddqwjgdvaujhyqsqtoexmnfuluaqbxoofvotvfoiexbnprrxptchmlctzgqtkivsilwgwgvpidpvasurraqfkcmxhdapjrlrnkbklwkrvoaziznlpor"
	t := "qhxepbshlrhoecdaodgpousbzfcqjxulatciapuftffahhlmxbufgjuxstfjvljybfxnenlacmjqoymvamphpxnolwijwcecgwbcjhgdybfffwoygikvoecdggplfohemfypxfsvdrseyhmvkoovxhdvoavsqqbrsqrkqhbtmgwaurgisloqjixfwfvwtszcxwktkwesaxsmhsvlitegrlzkvfqoiiwxbzskzoewbkxtphapavbyvhzvgrrfriddnsrftfowhdanvhjvurhljmpxvpddxmzfgwwpkjrfgqptrmumoemhfpojnxzwlrxkcafvbhlwrapubhveattfifsmiounhqusvhywnxhwrgamgnesxmzliyzisqrwvkiyderyotxhwspqrrkeczjysfujvovsfcfouykcqyjoobfdgnlswfzjmyucaxuaslzwfnetekymrwbvponiaojdqnbmboldvvitamntwnyaeppjaohwkrisrlrgwcjqqgxeqerjrbapfzurcwxhcwzugcgnirkkrxdthtbmdqgvqxilllrsbwjhwqszrjtzyetwubdrlyakzxcveufvhqugyawvkivwonvmrgnchkzdysngqdibhkyboyftxcvvjoggecjsajbuqkjjxfvynrjsnvtfvgpgveycxidhhfauvjovmnbqgoxsafknluyimkczykwdgvqwlvvgdmufxdypwnajkncoynqticfetcdafvtqszuwfmrdggifokwmkgzuxnhncmnsstffqpqbplypapctctfhqpihavligbrutxmmygiyaklqtakdidvnvrjfteazeqmbgklrgrorudayokxptswwkcircwuhcavhdparjfkjypkyxhbgwxbkvpvrtzjaetahmxevmkhdfyidhrdeejapfbafwmdqjqszwnwzgclitdhlnkaiyldwkwwzvhyorgbysyjbxsspnjdewjxbhpsvj"
	fmt.Printf("%s\n", string(findTheDifference2(s, t)))
	fmt.Printf("%s\n", string(findTheDifference3(s, t)))
}

func findTheDifference(s string, t string) byte {

	maps := make(map[byte]int)

	for i := range s {
		if _, ok := maps[s[i]]; !ok {
			maps[s[i]] = 1
		} else {
			maps[s[i]] += 1
		}
	}

	for i := range t {
		if _, ok := maps[t[i]]; ok {
			maps[t[i]] -= 1
			if maps[t[i]] < 0 {
				return t[i]
			}
		} else {
			return t[i]
		}
	}
	return byte(0)
}

func findTheDifference2(s string, t string) byte {
	ss := []byte(s)
	sort.Slice(ss, func(i, j int) bool {
		return ss[i] <= ss[j]
	})

	st := []byte(t)
	sort.Slice([]byte(st), func(i, j int) bool {
		return st[i] <= st[j]
	})

	for i := range ss {
		if ss[i] != st[i] {
			return st[i]
		}
	}
	return st[len(st)-1]

}

// 求和
func findTheDifference3(s string, t string) byte {

	sum1, sum2 := 0, 0
	for i := range s {
		sum1 += int(s[i])
		sum2 += int(t[i])
	}
	sum2 += int(t[len(t)-1])

	return byte(sum2 - sum1)
}
