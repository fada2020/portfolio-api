# 🚀 **Supabase 무료 배포 가이드**

Supabase를 사용하면 **PostgreSQL + 인증 + 실시간 + 스토리지**를 모두 무료로 사용할 수 있습니다!

## 🎯 **Supabase 무료 플랜**

### ✅ **무료로 제공되는 것들**
- **PostgreSQL**: 500MB 데이터베이스
- **인증**: 무제한 사용자 (소셜 로그인 포함)
- **실시간**: WebSocket 무료
- **스토리지**: 1GB 파일 저장
- **API**: 자동 생성 REST API
- **대시보드**: 웹 관리 인터페이스

### 📊 **제한사항**
- **API 요청**: 500MB 전송량/월 (충분함)
- **동시 연결**: 60개 (개인 포트폴리오로 충분)
- **행 수준 보안**: 무료 사용 가능

---

## 🔧 **1단계: Supabase 프로젝트 생성**

1. **Supabase 가입**: https://supabase.com 에서 무료 계정 생성
2. **새 프로젝트 생성**:
   - Organization: Personal
   - Project Name: `portfolio-api`
   - Database Password: 강력한 비밀번호 생성
   - Region: `Northeast Asia (Seoul)` (한국 사용자용)

3. **프로젝트 생성 완료** (약 2분 소요)

---

## 🔑 **2단계: 환경변수 설정**

프로젝트 설정에서 다음 정보를 확인:

```env
# Supabase 설정
SUPABASE_URL=https://your-project.supabase.co
SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
SUPABASE_SERVICE_ROLE_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...

# 데이터베이스 직접 연결
SUPABASE_DB_URL=postgresql://postgres:password@db.project.supabase.co:5432/postgres

# 기본 설정
GIN_MODE=release
PORT=8080
```

### 📍 **환경변수 위치 찾기**
1. Supabase 대시보드 → Settings → API
2. **Project URL**: `SUPABASE_URL`
3. **Project API keys**: `SUPABASE_ANON_KEY`
4. **Database** → Settings → Connection string

---

## 🛠️ **3단계: Railway에 배포**

### Railway에서 환경변수 설정:

1. **Railway 프로젝트** 생성 후 Variables 탭
2. **환경변수 추가**:
   ```
   SUPABASE_URL=https://your-project.supabase.co
   SUPABASE_ANON_KEY=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...
   SUPABASE_DB_URL=postgresql://postgres:password@db.project.supabase.co:5432/postgres
   GIN_MODE=release
   ```

3. **자동 배포** 시작!

---

## 📊 **4단계: 테이블 자동 생성 확인**

배포 후 Supabase 대시보드에서 확인:

1. **Table Editor** 탭 → 다음 테이블들이 자동 생성됨:
   - `users` - 사용자 정보
   - `projects` - 포트폴리오 프로젝트
   - `skills` - 기술 스택
   - `contact_messages` - 연락처 메시지
   - `analytics` - 방문 통계

2. **샘플 데이터** 자동 삽입 확인

---

## 🔐 **5단계: Row Level Security (RLS) 활용**

Supabase의 강력한 보안 기능 자동 설정:

```sql
-- 공개 프로필은 모든 사용자가 조회 가능
CREATE POLICY "Public profiles are viewable by everyone"
ON users FOR SELECT
USING (is_public = true);

-- 공개 프로젝트는 모든 사용자가 조회 가능
CREATE POLICY "Public projects are viewable by everyone"
ON projects FOR SELECT
USING (is_public = true);
```

---

## 🚀 **6단계: API 테스트**

배포 후 API 엔드포인트 테스트:

```bash
# 헬스체크
curl https://your-app.up.railway.app/health

# 사용자 조회
curl https://your-app.up.railway.app/api/v1/users

# 프로젝트 조회
curl https://your-app.up.railway.app/api/v1/projects

# 기술 스택 조회
curl https://your-app.up.railway.app/api/v1/skills
```

---

## 📈 **7단계: Supabase 대시보드 활용**

### 실시간 모니터링:
1. **API Logs**: 모든 요청 로그 확인
2. **Database**: 실시간 데이터 확인/수정
3. **Authentication**: 사용자 관리 (추후 확장 시)
4. **Storage**: 이미지 업로드 기능 (추후 확장 시)

### 데이터 관리:
- **Table Editor**: 직접 데이터 추가/수정
- **SQL Editor**: 복잡한 쿼리 실행
- **API Docs**: 자동 생성된 API 문서

---

## 🔥 **Supabase의 추가 장점**

### 🎨 **향후 확장 가능한 기능들**
1. **소셜 로그인**: GitHub, Google 등 무료
2. **실시간 업데이트**: WebSocket으로 실시간 데이터
3. **파일 업로드**: 프로젝트 이미지 저장
4. **이메일 인증**: 자동 이메일 발송
5. **Edge Functions**: 서버리스 함수 (Beta)

### 📊 **무료 사용량 모니터링**
- Dashboard → Settings → Usage
- API 요청, 데이터베이스 크기, 스토리지 사용량 실시간 확인

---

## 💡 **최적화 팁**

### 1. **쿼리 최적화**
```sql
-- 인덱스 활용 (이미 생성됨)
CREATE INDEX idx_projects_featured ON projects(featured);
CREATE INDEX idx_analytics_page ON analytics(page);
```

### 2. **API 캐싱**
```go
// Redis 없이도 메모리 캐싱 가능
var cache = make(map[string]interface{})
```

### 3. **이미지 최적화**
- Supabase Storage 사용 시 자동 리사이징
- WebP 형식 자동 변환

---

## 🆘 **문제 해결**

### 연결 실패 시:
```bash
# 연결 테스트
psql "postgresql://postgres:password@db.project.supabase.co:5432/postgres"
```

### RLS 정책 확인:
```sql
-- Supabase SQL Editor에서 실행
SELECT * FROM pg_policies WHERE tablename = 'users';
```

### API 로그 확인:
- Supabase Dashboard → Logs → API Logs

---

## 🎉 **완료!**

이제 **Supabase + Railway**로 완전 무료 포트폴리오 API가 완성되었습니다!

- **데이터베이스**: Supabase PostgreSQL (무료)
- **호스팅**: Railway (무료)
- **도메인**: 무료 서브도메인
- **모니터링**: Supabase 대시보드
- **확장성**: 필요 시 유료로 쉽게 업그레이드

**총 비용: $0/월** 🎯