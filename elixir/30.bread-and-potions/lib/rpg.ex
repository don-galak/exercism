defmodule RPG do
  defmodule Character do
    defstruct health: 100, mana: 0
  end

  defmodule LoafOfBread do
    defstruct []
  end

  defmodule ManaPotion do
    defstruct strength: 10
  end

  defmodule Poison do
    defstruct []
  end

  defmodule EmptyBottle do
    defstruct []
  end

  defprotocol Edible do
    def eat(item, character)
  end

  defimpl Edible, for: LoafOfBread do
    @bread_health 5
    def eat(_, %{health: health} = character) do
      {nil, %{character | health: health + @bread_health}}
    end
  end

  defimpl Edible, for: ManaPotion do
    def eat(%{strength: strength}, %{mana: mana} = character) do
      {%EmptyBottle{}, %{character | mana: mana + strength}}
    end
  end

  defimpl Edible, for: Poison do
    def eat(_, character) do
      {%EmptyBottle{}, %{character | health: 0}}
    end
  end
end
