<script lang="ts">
  // The HireMe brand mark — an "H" whose right stem rises into an upward arrow:
  // get hired, move up. `fill="currentColor"` makes it track the theme (dark mark
  // on light, light mark on dark), exactly like the wordmark beside it. Shared by
  // the desktop header (TopBar) and the mobile drawer (HeaderMenu) so the geometry
  // lives in one place. aria-hidden: callers pair it with the visible "HireMe"
  // text, which names the link. The OG-image copy in server/og/shared.ts is a
  // separately rendered PNG (satori) and intentionally does not route through here.
  let { class: cls = 'size-5 shrink-0' }: { class?: string } = $props();
</script>

<svg
  viewBox="0 0 512 512"
  class={cls}
  fill="currentColor"
  aria-hidden="true"
  xmlns="http://www.w3.org/2000/svg"
>
  <g>
    <rect x="172" y="168" width="44" height="176" />
    <rect x="296" y="128" width="44" height="216" />
    <rect x="172" y="234" width="168" height="40" />
    <path d="M296 128 L340 128 L318 82 Z" />
  </g>
</svg>
