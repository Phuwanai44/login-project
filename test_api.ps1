$baseUrl = 'http://localhost:3000'

Write-Host ('=' * 50) -ForegroundColor Cyan
Write-Host '  SYSTEM API TEST SUITE' -ForegroundColor Cyan
Write-Host ('=' * 50) -ForegroundColor Cyan

# ─── TEST 1: REGISTER ───────────────────────────────
Write-Host ''
Write-Host '[TEST 1] POST /register' -ForegroundColor Yellow
try {
    $body = '{"name":"Test User","email":"testuser88@example.com","password":"testpass123"}'
    $res = Invoke-RestMethod -Uri "$baseUrl/register" -Method POST -ContentType 'application/json' -Body $body -ErrorAction Stop
    Write-Host '  PASS - Register successful' -ForegroundColor Green
    Write-Host ("  Response: " + ($res | ConvertTo-Json -Compress))
} catch {
    $code = $_.Exception.Response.StatusCode.value__
    Write-Host ("  INFO - Status $code : " + $_.ErrorDetails.Message) -ForegroundColor Magenta
}

# ─── TEST 2: LOGIN ──────────────────────────────────
Write-Host ''
Write-Host '[TEST 2] POST /login' -ForegroundColor Yellow
$token = ''
$userRole = ''
try {
    $body = '{"email":"testuser88@example.com","password":"testpass123"}'
    $res = Invoke-RestMethod -Uri "$baseUrl/login" -Method POST -ContentType 'application/json' -Body $body -ErrorAction Stop
    $token = $res.token
    $userRole = $res.user.role
    $tokenPreview = $token.Substring(0, [Math]::Min(40, $token.Length))
    Write-Host '  PASS - Login successful' -ForegroundColor Green
    Write-Host "  Token: $tokenPreview..."
    Write-Host "  Role:  $userRole"
} catch {
    Write-Host ('  FAIL - ' + $_.ErrorDetails.Message) -ForegroundColor Red
}

# ─── TEST 3: GET PROFILE ────────────────────────────
Write-Host ''
Write-Host '[TEST 3] GET /profile (auth required)' -ForegroundColor Yellow
if ($token -ne '') {
    try {
        $headers = @{ Authorization = "Bearer $token" }
        $res = Invoke-RestMethod -Uri "$baseUrl/profile" -Method GET -Headers $headers -ErrorAction Stop
        Write-Host '  PASS - Profile fetched' -ForegroundColor Green
        Write-Host ("  Name: " + $res.name + " | Email: " + $res.email + " | Role: " + $res.role)
    } catch {
        Write-Host ('  FAIL - ' + $_.ErrorDetails.Message) -ForegroundColor Red
    }
} else {
    Write-Host '  SKIP - No token available' -ForegroundColor DarkGray
}

# ─── TEST 4: GET USERS (admin) ──────────────────────
Write-Host ''
Write-Host '[TEST 4] GET /admin/users (admin role required)' -ForegroundColor Yellow
if ($token -ne '') {
    try {
        $headers = @{ Authorization = "Bearer $token" }
        $res = Invoke-RestMethod -Uri "$baseUrl/admin/users" -Method GET -Headers $headers -ErrorAction Stop
        Write-Host ("  PASS - Got " + $res.Count + " users") -ForegroundColor Green
    } catch {
        $code = $_.Exception.Response.StatusCode.value__
        Write-Host ("  INFO - Status $code (expected if non-admin user)") -ForegroundColor Magenta
    }
} else {
    Write-Host '  SKIP - No token available' -ForegroundColor DarkGray
}

# ─── TEST 5: GET PRODUCTS ───────────────────────────
Write-Host ''
Write-Host '[TEST 5] GET /products' -ForegroundColor Yellow
if ($token -ne '') {
    try {
        $headers = @{ Authorization = "Bearer $token" }
        $res = Invoke-RestMethod -Uri "$baseUrl/products" -Method GET -Headers $headers -ErrorAction Stop
        Write-Host '  PASS - Products fetched' -ForegroundColor Green
        $json = $res | ConvertTo-Json -Compress
        $preview = $json.Substring(0, [Math]::Min(120, $json.Length))
        Write-Host "  Response: $preview..."
    } catch {
        Write-Host ('  FAIL - ' + $_.ErrorDetails.Message) -ForegroundColor Red
    }
} else {
    Write-Host '  SKIP - No token available' -ForegroundColor DarkGray
}

# ─── TEST 6: PROTECTED ROUTE WITHOUT TOKEN ──────────
Write-Host ''
Write-Host '[TEST 6] GET /profile without token (expect 401)' -ForegroundColor Yellow
try {
    $res = Invoke-RestMethod -Uri "$baseUrl/profile" -Method GET -ErrorAction Stop
    Write-Host '  FAIL - Should have been blocked!' -ForegroundColor Red
} catch {
    $code = $_.Exception.Response.StatusCode.value__
    if ($code -eq 401) {
        Write-Host '  PASS - Correctly blocked with 401 Unauthorized' -ForegroundColor Green
    } else {
        Write-Host ("  INFO - Got status $code") -ForegroundColor Magenta
    }
}

Write-Host ''
Write-Host ('=' * 50) -ForegroundColor Cyan
Write-Host '  TEST SUITE COMPLETE' -ForegroundColor Cyan
Write-Host ('=' * 50) -ForegroundColor Cyan
