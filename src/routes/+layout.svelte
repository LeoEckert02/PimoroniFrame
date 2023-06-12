<script>
    import "../app.css";
    import {Toaster} from 'svelte-french-toast';
</script>

<Toaster/>
<body class="min-h-screen bg-custom1">
    <div class="bg-dotted-spacing-3 bg-dotted-amber-300 min-h-screen">
    <slot/>
    </div>
</body>

