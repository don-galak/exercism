defmodule Newsletter do
  @spec read_emails(path :: String.t()) :: list(String.t())
  def read_emails(path), do: String.split(File.read!(path))

  def open_log(path), do: File.open!(path, [:write])

  def log_sent_email(pid, email), do: IO.puts(pid, email)

  def close_log(pid), do: File.close(pid)

  @spec send_newsletter(String.t(), pid(), (String.t() -> any())) :: any
  def send_newsletter(emails_path, log_path, send_fun) do
    pid = open_log(log_path)

    Enum.each(read_emails(emails_path), fn email ->
      if send_fun.(email) == :ok, do: log_sent_email(pid, email)
    end)

    close_log(pid)
  end
end
