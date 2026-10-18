<script lang="ts">
  import { resolve } from '$app/paths';
  import { popularCollectionLinks } from '$lib/collections';
  import { reopen } from '$lib/consent.svelte';
  import { EXTENSION_STORE_URL } from '$lib/extensionLinks';
  import { SOCIAL_LINKS } from '$lib/socialLinks';
  import { ProviderIcon } from '$lib/ui';

  // Grouped navigation over existing routes only — kept deliberately small so the
  // footer stays uncluttered. Internal links go through resolve() (base-path safe),
  // mirroring the header.
  const groups = [
    {
      title: 'Product',
      links: [
        { label: 'Jobs', href: resolve('/jobs') },
        { label: 'Companies', href: resolve('/companies') },
        { label: 'Collections', href: resolve('/collections') },
        { label: 'Talent Network', href: resolve('/talent') },
        { label: 'Jobs by role', href: resolve('/roles') },
        { label: 'Pricing', href: resolve('/pricing') },
        // The glossary's only link from the app — the chip's reveal opens on
        // interaction and the sitemap is for crawlers, so this is the one path a
        // reader browsing the site can follow to it.
        { label: 'Skills glossary', href: resolve('/skills') },
        { label: 'Recruiters', href: resolve('/recruiters') },
      ],
    },
    {
      title: 'Features',
      links: [
        { label: 'Advanced search', href: resolve('/features/advanced-search') },
        { label: 'Browser extension', href: resolve('/features/extension') },
        { label: 'Inbox', href: resolve('/features/inbox') },
        { label: 'CV tailoring', href: resolve('/features/tailor') },
        { label: 'Application tracking', href: resolve('/features/tracking') },
        { label: 'Notifications', href: resolve('/features/notifications') },
        { label: 'Referrals', href: resolve('/features/referrals') },
        { label: 'Ghost jobs', href: resolve('/features/ghost-jobs') },
      ],
    },
    {
      title: 'Resources',
      links: [
        { label: 'How it works', href: resolve('/how-it-works') },
        { label: 'Discussions', href: resolve('/discussions') },
        { label: 'Blog', href: resolve('/blog') },
        { label: 'Insights', href: resolve('/insights') },
        { label: 'Hiring signal', href: resolve('/insights/companies') },
        { label: 'Trends', href: resolve('/trends') },
        { label: 'For AI agents', href: resolve('/agents') },
        { label: 'CLI', href: resolve('/cli') },
        { label: 'ChatGPT', href: resolve('/chatgpt') },
        { label: 'API docs', href: resolve('/docs/api') },
      ],
    },
    {
      title: 'Company',
      links: [
        { label: 'About', href: resolve('/about') },
        { label: 'Open', href: resolve('/open') },
        { label: 'For companies', href: resolve('/for-companies') },
        { label: 'Contribute', href: resolve('/contribute') },
        // Next to Contribute rather than under Resources: the two are one thought —
        // here is how to help, and here is everyone who did.
        { label: 'Contributors', href: resolve('/contributors') },
        { label: 'Submit a job', href: resolve('/submit') },
        // Next to Status rather than under Resources: the two answer neighbouring
        // questions — where the jobs come from, and whether we are still reading them.
        { label: 'Sources', href: resolve('/sources') },
        { label: 'Status', href: resolve('/status') },
        { label: 'Support', href: resolve('/support') },
        { label: 'Privacy', href: resolve('/privacy') },
        { label: 'Terms', href: resolve('/terms') },
      ],
    },
  ];

  // External profiles: open in a new tab, each rendered with its ProviderIcon brand
  // mark. All follow the muted text colour (so they match and hover works). Shared
  // with the /signin brand panel — see $lib/socialLinks.

  // A strip below the four groups rather than a fifth column: the grid is
  // sm:grid-cols-4, and ten links would not fit one anyway. Kept out of `groups`
  // because these are collection landing pages, not site navigation — and because
  // it is the one place every page links them from (see popularCollectionLinks).
  const popular = popularCollectionLinks();

  const year = new Date().getFullYear();

  // The extension's source, next to the link columns: a download button among
  // page links reads as neither, so it gets its own row.
  const stores = [
    {
      provider: 'github',
      kicker: 'Open source',
      name: 'Extension source',
      href: EXTENSION_STORE_URL,
    },
  ];

  // Compact: the account shell (/my/*) is an app-like surface with its own sidebar
  // nav, so the four link columns, the popular-collections strip and the Product
  // Hunt badge below are marketing chrome that doesn't belong there — only the
  // bottom bar (copyright, cookie settings, social links, open-source note) still
  // applies. Full-bleed account routes (/my/assistant/*, /tailor/*) render no
  // footer at all, compact or otherwise — see routes/+layout.svelte's hideFooter.
  let { compact = false }: { compact?: boolean } = $props();
</script>

<footer class="border-t border-border">
  {#if !compact}
    <div class="mx-auto max-w-6xl px-4 py-8 sm:py-12">
      <div class="grid grid-cols-2 gap-x-6 gap-y-7 sm:grid-cols-4 sm:gap-6">
        <!-- Navigation groups. Each is a named landmark (aria-label) so screen readers
             get a title without adding headings to the page outline. -->
        {#each groups as group (group.title)}
          <nav class="flex flex-col gap-3" aria-label={group.title}>
            <p class="text-xs font-medium uppercase tracking-wider text-muted-foreground">
              {group.title}
            </p>
            <ul class="flex flex-col gap-2">
              {#each group.links as link (link.href)}
                <li>
                  <!-- eslint-disable-next-line svelte/no-navigation-without-resolve -- internal route already passed through resolve() when building `groups`; the linter can't trace it via the variable -->
                  <a href={link.href}
                    class="text-sm text-muted-foreground transition-colors hover:text-foreground"
                  >
                    {link.label}
                  </a>
                </li>
              {/each}
            </ul>
          </nav>
        {/each}
      </div>

      <!-- Popular collections. Real <a href> in the server-rendered HTML: crawlers
           discover links by parsing markup, and these landing pages had none from the
           homepage at all. -->
      <nav class="mt-8 border-t border-border pt-6" aria-label="Popular collections">
        <p class="text-xs font-medium uppercase tracking-wider text-muted-foreground">Popular</p>
        <ul class="mt-3 flex flex-wrap gap-x-4 gap-y-2">
          {#each popular as collection (collection.slug)}
            <li>
              <a href={resolve('/collections/[slug]', { slug: collection.slug })}
                class="text-sm text-muted-foreground transition-colors hover:text-foreground"
              >
                {collection.title}
              </a>
            </li>
          {/each}
        </ul>
      </nav>

      <!-- The extension source button gets its own row, centred: a download
           button among page links reads as neither. Wrapping, so a narrow viewport
           stacks rather than shrinks. -->
      <div class="mt-8 flex flex-wrap items-center gap-3">
        {#each stores as store (store.href)}
          <!-- eslint-disable-next-line svelte/no-navigation-without-resolve -- external repository URL opened in a new tab; not an internal route -->
          <a href={store.href}
            target="_blank"
            rel="noopener noreferrer"
            class="inline-flex h-14 items-center gap-3 rounded-lg border border-border px-4 text-foreground transition-colors hover:bg-muted"
          >
            <ProviderIcon provider={store.provider} class="size-7 shrink-0" />
            <span class="flex flex-col leading-tight">
              <span class="text-xs uppercase tracking-wider text-muted-foreground">
                {store.kicker}
              </span>
              <span class="text-sm font-semibold">{store.name}</span>
            </span>
          </a>
        {/each}
      </div>
    </div>
  {/if}

  <!-- Bottom bar: copyright + social links on the left, open-source note on the
       right. Split off by its own top border only when it follows the link
       groups above — compact, the footer's own top border already does that job. -->
  <div class={compact ? undefined : 'border-t border-border'}>
    <div
      class="mx-auto flex max-w-6xl flex-col gap-3 px-4 py-4 text-xs text-muted-foreground sm:flex-row sm:items-center sm:justify-between sm:gap-1"
    >
      <div class="flex items-center gap-4">
        <p>© {year}</p>
        <!-- Re-opens the consent banner so a prior cookie choice can be changed —
             withdrawal as easy as granting (GDPR). -->
        <button
          type="button"
          onclick={reopen}
          class="text-muted-foreground transition-colors hover:text-foreground"
        >
          Cookie settings
        </button>
        <div class="flex items-center gap-3">
          {#each SOCIAL_LINKS as social (social.provider)}
            <!-- eslint-disable-next-line svelte/no-navigation-without-resolve -- external profile URL opened in a new tab; not an internal route -->
            <a href={social.href}
              target="_blank"
              rel="noopener noreferrer"
              aria-label={social.label}
              class="text-muted-foreground transition-colors hover:text-foreground"
            >
              <ProviderIcon provider={social.provider} />
            </a>
          {/each}
        </div>
      </div>
      <p>
        Free &amp; open-source.
        <!-- eslint-disable-next-line svelte/no-navigation-without-resolve -- external repository URL opened in a new tab; not an internal route -->
        <a href="https://github.com/dunsin2018/HireMe"
          target="_blank"
          rel="noopener noreferrer"
          class="font-medium text-foreground transition-colors hover:text-muted-foreground"
        >
          View source on GitHub
        </a>.
      </p>
    </div>
  </div>
</footer>
