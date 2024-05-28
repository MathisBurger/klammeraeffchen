<script>
	let loginBaseUrl = process.env.NODE_ENV === "production" ? "/login" : "http://localhost:3000/login"
</script>

<svelte:head>
	<title>Klammeraeffchen</title>
</svelte:head>

<section>
	<h1>Landing page</h1>
	<a href="{loginBaseUrl}">Login</a>
</section>
