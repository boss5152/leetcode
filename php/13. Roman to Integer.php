class Solution {
    /**
     * @param String $s
     * @return Integer
     */
    function romanToInt($s) {
        // 切割成陣列
        $aRoman = str_split($s);
        
        return $this->romanConvertInt($aRoman);
    }
    
    /**
     * 羅馬轉數字
     *
     * @param array $aRoman 羅馬文字陣列
     * @param string $nowKey 當前
     * @param string $nextKey 下一個
     * @param int $answer 答案
     *
     * @return int
     **/
    private function romanConvertInt($aRoman, $nowKey = 0, $nextKey = 1, $answer = 0)
    {
        // 當前羅馬文字
        $now = $aRoman[$nowKey];
        
        // 下一個羅馬文字
        $next = $aRoman[$nextKey] ?? 'end';

        switch ($now) {
            case 'I':
                $answer += (in_array($next, ['V', 'X'])) ? -1 : 1;
                break;
            case 'V':
                $answer += 5;
                break;
            case 'X':
                $answer += (in_array($next, ['L', 'C'])) ? -10 : 10;
                break;
            case 'L':
                $answer += 50;
                break;
            case 'C':
                $answer += (in_array($next, ['D', 'M'])) ? -100 : 100;
                break;
            case 'D':
                $answer += 500;
                break;
            case 'M':
                $answer += 1000;
                break;
        }
        
        // 還沒結束，繼續加
        if ($next !== 'end') {
            // 遞迴
            $answer = $this->romanConvertInt($aRoman, $nextKey, $nextKey + 1, $answer);
        }

        return $answer;
    }
}