# yabaip

Like tmuxp, but for [yabai](https://github.com/koekeishiya/yabai). Also, it doesn't do anything yet.

Just a mini-project to try to learn Golang; please be gentle. 😅

## Spaces

```yaml
label: my_space
windows: 
  - 
    command: 
      - open
      - /Users/prescott
    window_type: managed
  - 
    command: 
      - open
      - /Applications/kitty.app
    layout: 
      window_type: floating
```