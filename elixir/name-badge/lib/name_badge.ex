defmodule NameBadge do
  def print(id, name, department) do
    dep = if department, do: String.upcase(department), else: "OWNER"
    employee_id = if id, do: "[#{id}] - ", else: ""
    "#{employee_id}#{name} - #{dep}"
  end
end
