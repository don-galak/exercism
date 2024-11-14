defmodule Newsletter do
  def read_emails(path) do
    {_, contents} = File.read(path)
    if String.length(contents) > 0, do: String.split(String.trim(contents), "\n"), else: []
  end

  def open_log(path) do
    {_, pid} = File.open(path, [:write])
    pid
  end

  def log_sent_email(pid, email), do: IO.puts(pid, email)

  def close_log(pid), do: File.close(pid)

  def send_newsletter(emails_path, log_path, send_fun) do
    # Please implement the send_newsletter/3 function
  end
end
