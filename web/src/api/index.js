import api from './config'

export function getUser () {
  return api.post('/test', { name: 'young' })
}
