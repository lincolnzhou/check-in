import http from '../request'
import { uri_check }from '../common/port_uri'

export function getHitCount() {
	return http.get(uri_check.hit_count)
}

export function getCheckList() {
	return http.get(uri_check.check_list)
}

export function postCheck() {
	return http.post(uri_check.post_check)
}
