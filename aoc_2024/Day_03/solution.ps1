$corruptedMemory = Get-Content -Path "input"

$patternMul = "mul\((\d+),(\d+)\)"
$totalSumPart1 = 0

$matches = [regex]::Matches($corruptedMemory, $patternMul)

foreach ($match in $matches) {
    $x = [int]$match.Groups[1].Value
    $y = [int]$match.Groups[2].Value

    $totalSumPart1 += $x * $y
}

$patternDo = "do\(\)"
$patternDont = "don't\(\)"
$isEnabled = $true
$totalSumPart2 = 0

$instructions = [regex]::Split($corruptedMemory, "(?=mul\(|do\(\)|don't\(\))")

foreach ($instruction in $instructions) {
    if ($instruction -match $patternDo) {
        $isEnabled = $true
    } elseif ($instruction -match $patternDont) {
        $isEnabled = $false
    } elseif ($instruction -match $patternMul) {
        if ($isEnabled) {
            $mulMatch = [regex]::Match($instruction, $patternMul)
            if ($mulMatch.Success) {
                $x = [int]$mulMatch.Groups[1].Value
                $y = [int]$mulMatch.Groups[2].Value
                $totalSumPart2 += $x * $y
            }
        }
    }
}

Write-Output "Part 1: The sum of all valid mul(X,Y) results is: $totalSumPart1"
Write-Output "Part 2: The sum of all enabled mul(X,Y) results is: $totalSumPart2"

