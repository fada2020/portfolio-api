# 🔒 **보안 환경변수 설정 가이드**

## 📍 **Step 1: 로컬 환경변수 설정**

### 1.1 .env 파일 생성 (Git에 올리지 않음)
```bash
cd /path/to/portfolio-api
touch .env
```

### 1.2 .env 파일에 실제 값 입력
```bash
# .env 파일 내용 (절대 Git에 올리지 마세요!)
SUPABASE_URL=https://abcdefghijklmnop.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.실제키값
SUPABASE_DB_URL=postgresql://postgres:실제비밀번호@db.abcdefghijklmnop.supabase.co:5432/postgres
GIN_MODE=release
```

### 1.3 .gitignore 확인
```bash
# .gitignore에 이미 추가됨
.env
.env.local
.env.production
```

---

## 📍 **Step 2: Git 안전 업로드**

### 2.1 안전한 파일들만 커밋
```bash
git add .
git status
```

확인사항:
- ✅ `.env.example` - 예시 파일 (안전)
- ✅ `.gitignore` - 보안 설정 (안전)
- ✅ `vercel.json` - 배포 설정 (안전)
- ❌ `.env` - 실제 키값 (Git에 없어야 함)

### 2.2 커밋 및 푸시
```bash
git commit -m "🔒 Add security configs and deployment files"
git push origin main
```

---

## 📍 **Step 3: Vercel 환경변수 설정**

### 3.1 Supabase에서 키값 복사

**Supabase Dashboard → Settings → API**:
- `Project URL` 복사
- `anon public` 키 복사

**Settings → Database → Connection string**:
- PostgreSQL URL 복사

### 3.2 Vercel CLI로 안전하게 설정
```bash
# 1. SUPABASE_URL 설정
vercel env add SUPABASE_URL
? What's the value of SUPABASE_URL? https://your-actual-project.supabase.co
? Add SUPABASE_URL to which Environments? Production, Preview, Development

# 2. SUPABASE_ANON_KEY 설정
vercel env add SUPABASE_ANON_KEY
? What's the value of SUPABASE_ANON_KEY? eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...실제키
? Add SUPABASE_ANON_KEY to which Environments? Production, Preview, Development

# 3. SUPABASE_DB_URL 설정
vercel env add SUPABASE_DB_URL
? What's the value of SUPABASE_DB_URL? postgresql://postgres:실제비밀번호@db...
? Add SUPABASE_DB_URL to which Environments? Production, Preview, Development

# 4. GIN_MODE 설정
vercel env add GIN_MODE
? What's the value of GIN_MODE? release
? Add GIN_MODE to which Environments? Production, Preview, Development
```

### 3.3 환경변수 확인
```bash
vercel env ls
```

---

## 📍 **Step 4: 배포 및 테스트**

### 4.1 Vercel 배포
```bash
vercel --prod
```

### 4.2 환경변수 로딩 확인
```bash
curl https://your-app.vercel.app/health
```

성공 응답:
```json
{
  "status": "healthy",
  "timestamp": "2024-09-26T...",
  "version": "1.0.0",
  "platform": "vercel"
}
```

---

## 🛡️ **보안 체크포인트**

### ✅ **안전 확인사항**
- [ ] `.env` 파일이 Git에 추가되지 않음
- [ ] `.gitignore`에 `.env` 포함됨
- [ ] GitHub에서 `.env` 파일이 보이지 않음
- [ ] Vercel 환경변수가 CLI로 설정됨
- [ ] API가 정상 동작함

### ❌ **위험 신호**
- GitHub에서 실제 키값이 보임
- `.env` 파일이 커밋됨
- 하드코딩된 비밀키가 있음

---

## 💡 **추가 보안 팁**

### 1. **Supabase RLS 활용**
- Row Level Security로 데이터 접근 제한
- PUBLIC 테이블만 익명 접근 허용

### 2. **API 키 순환**
- 정기적으로 Supabase 키 재생성
- Vercel 환경변수 업데이트

### 3. **모니터링**
- Supabase Dashboard에서 API 사용량 확인
- 비정상적인 접근 패턴 모니터링

---

## 🔥 **응급 상황 대처**

### Git에 실수로 키를 올렸다면:
```bash
# 1. 즉시 Supabase에서 키 재생성
# 2. Git 히스토리 정리
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch .env' --prune-empty --tag-name-filter cat -- --all

# 3. 강제 푸시
git push origin main --force

# 4. 새 키로 Vercel 환경변수 업데이트
vercel env rm SUPABASE_ANON_KEY
vercel env add SUPABASE_ANON_KEY
```