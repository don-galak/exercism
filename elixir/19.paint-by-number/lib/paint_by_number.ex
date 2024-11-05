defmodule PaintByNumber do
  def palette_bit_size(color_count, n \\ 0) do
    cond do
      Integer.pow(2, n) >= color_count -> n
      true -> palette_bit_size(color_count, n + 1)
    end
  end

  def empty_picture(), do: <<>>

  def test_picture(), do: <<0::2, 1::2, 2::2, 3::2>>

  def prepend_pixel(picture, color_count, pixel_color_index) do
    # Please implement the prepend_pixel/3 function
  end

  def get_first_pixel(picture, color_count) do
    # Please implement the get_first_pixel/2 function
  end

  def drop_first_pixel(picture, color_count) do
    # Please implement the drop_first_pixel/2 function
  end

  def concat_pictures(picture1, picture2) do
    # Please implement the concat_pictures/2 function
  end
end

# 001
# 010
# 011
# 110
# 111
