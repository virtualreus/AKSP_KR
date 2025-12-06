export interface HealthMetric {
  id: string
  user_id?: string
  metric_type: string
  value: number
  unit: string
  recorded_at: string
  notes?: string
  created_at?: string
}

