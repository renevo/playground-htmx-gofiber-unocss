# Playground

This uses the following items from [A very simple tech stack](https://www.youtube.com/watch?v=huMTT5Pb8b8).

* [UnoCSS](https://unocss.dev/)
* [gofiber](https://github.com/gofiber/fiber)
* [htmx](https://htmx.org/)

## Polygon API

Application requires a [Polygon](https://polygon.io/docs/stocks/getting-started) API key, sign up for a free account and use the default API key set to 
the environment variable: `POLYGON_API_KEY`.

The free tier is limited to 5 requests per minute.

## Thoughts

It's ugly, but functional, and does what it says on paper. The video missed a lot of things I had to fill the gaps for, but overall, a very simple tech stack. Although, Unocss doesn't look very mature and for people who want to build there own entire CSS framework, [Tailwind](https://tailwindcss.com/) might be a better fit.

Fiber does some templates a bit different when including them [html.Engine](https://docs.gofiber.io/template/html/)
