defmodule FreelancerRates do
  def daily_rate(hourly_rate) do
    hourly_rate * 8.0
  end

  def apply_discount(before_discount, discount) do
    before_discount - before_discount * discount / 100
  end

  def monthly_rate(hourly_rate, discount) do
    rate = daily_rate(hourly_rate) * 22
    trunc(ceil(apply_discount(rate, discount)))
  end

  def days_in_budget(budget, hourly_rate, discount) do
    Float.floor(22 * budget / monthly_rate(hourly_rate, discount), 1)
  end
end
