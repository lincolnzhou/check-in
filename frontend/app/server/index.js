import http from '../request/index.js'

export function getHitCount() {
	return http.get("/api/hit_count")
}
