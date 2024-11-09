defmodule LibraryFees do
  def datetime_from_string(string) do
    NaiveDateTime.from_iso8601!(string)
  end

  def before_noon?(datetime) do
    datetime.hour < 12
  end

  def return_date(checkout_datetime) do
    if before_noon?(checkout_datetime) do
      NaiveDateTime.to_date(NaiveDateTime.add(checkout_datetime, 28, :day))
    else
      NaiveDateTime.to_date(NaiveDateTime.add(checkout_datetime, 29, :day))
    end
  end

  def days_late(planned_return_date, actual_return_datetime) do
    diff = Date.diff(actual_return_datetime, planned_return_date)
    if diff < 0 do
      0
    else
      diff
    end
  end

  def monday?(datetime) do
    Date.day_of_week(NaiveDateTime.to_date(datetime)) == 1
  end

  def calculate_late_fee(checkout, return, rate) do
    actual_return_datetime = datetime_from_string(return)
    checkout_date = return_date(datetime_from_string(checkout))

    fee = days_late(checkout_date, actual_return_datetime) * rate

    if monday?(actual_return_datetime) do
      trunc(fee / 2)
    else
      fee
    end
  end
end
