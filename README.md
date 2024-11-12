# Nadia's Webpage

This project is for Nadia. The code will be private, at least for now. I'm gonna make a great webpage.

## Pages

- home `"nadiatumiwa/"`
	- main image
	- navigation bar (sidebar preferably)
	- brief biography
		- text
		- link to full bio
	- performances
	- footer
		- social media
- biography `"nadiatumiwa/biography"`
- performances `"nadiatumiwa/performances"`
- gallery `"nadiatumiwa/photos"`
- videos `"nadiatumiwa/videos"`
- contact `"nadiatumiwa/contact"`
- admin `"nadiatumiwa/admin"`

## Features

- custom admin page
	- strong authentication (OAuth)
	- update dynamic content in db
		- add, edit, delete performances
		- update biography
		- upload/link photos & videos (?)
		- track basic analytics (?)
	- google sheets integration (?)
	- ensure the endpoint is not publicly discoverable
		- ensure search engines do not index the endpoint
	- have strong session management to prevent session hijacking
- frontend considerations
	- clean and engaging design
		- sidebar animations
	- prioritize mobile responsiveness
	- organize content clearly
	- fetch dynamic content from db
- database design
	- `performance` table
	- `biography` table
- performance & SEO
	- optimize images
	- use lazy loading for media galleries
	- include metadata for all pages
- contact
	- include a contact form (?)
	- display social media links prominently