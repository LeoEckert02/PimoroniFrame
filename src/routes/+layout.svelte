<script>
    import "../app.css";
    import { Toaster } from 'svelte-french-toast';
</script>

<Toaster />
<body class="min-h-screen bg-slate-50">
    <div class="flex items-center justify-center h-screen">
        <slot />
    </div>
</body>

