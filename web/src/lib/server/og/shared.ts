// Shared primitives for the Open Graph cards (job, company, brand). Pure and
// framework-free so every card speaks one brand vocabulary and cannot drift, and
// so the render smoke tests can exercise them directly.
//
// satori constraint: layout is flexbox only (no CSS grid), and any element with
// more than one child declares `display: flex`.

export const OG_WIDTH = 1200;
export const OG_HEIGHT = 630;

// The HireMe brand mark (an "H" whose right stem rises into an upward arrow),
// inlined as a base64 data-URI so satori embeds it with no network fetch (satori
// cannot fetch remote images). Coloured #0a0a0a to match the footer wordmark
// beside it.
const MARK_SVG =
  '<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">' +
  '<g fill="#0a0a0a">' +
  '<rect x="172" y="168" width="44" height="176"/>' +
  '<rect x="296" y="128" width="44" height="216"/>' +
  '<rect x="172" y="234" width="168" height="40"/>' +
  '<path d="M296 128 L340 128 L318 82 Z"/>' +
  '</g>' +
  '</svg>';
export const MARK_DATA_URI = `data:image/svg+xml;base64,${Buffer.from(MARK_SVG).toString('base64')}`;

export function esc(s: string): string {
  return s
    .replace(/&/g, '&amp;')
    .replace(/</g, '&lt;')
    .replace(/>/g, '&gt;')
    .replace(/"/g, '&quot;');
}

/** One or two uppercase initials from a name, for the logo fallback tile. */
function monogram(name: string): string {
  const words = name.trim().split(/\s+/).filter(Boolean);
  const first = words[0];
  if (!first) return '?';
  if (words.length === 1) return first.slice(0, 2).toUpperCase();
  return (first.charAt(0) + (words[1] ?? '').charAt(0)).toUpperCase();
}

/** A logo tile at `size`px: the resolved image, or a monogram tile derived from `name`
 *  when `logo` is null. Radius and monogram font scale with the size (at 72px this
 *  matches the original job-card tile: radius 14, font 30).
 *
 *  `shape` carries the same distinction the design system's Avatar makes: a rounded
 *  `tile` is an entity's mark, a `circle` is a person. A card about a person that
 *  squares them off reads as a company logo. */
export function logoBox(
  logo: string | null,
  name: string,
  size: number,
  shape: 'tile' | 'circle' = 'tile',
): string {
  const radius = shape === 'circle' ? '50%' : `${Math.round(size / 5)}px`;
  if (logo) {
    // satori reads image dimensions from style, not the width/height HTML attributes.
    // A person's avatar is cropped to fill its circle; an entity's mark is fitted
    // inside its tile, because a logo cropped at the edges is a broken logo.
    const fit = shape === 'circle' ? 'cover' : 'contain';
    return `<img src="${esc(logo)}" style="width:${size}px;height:${size}px;border-radius:${radius};object-fit:${fit}" />`;
  }
  const font = Math.round(size * 0.42);
  return (
    `<div style="display:flex;align-items:center;justify-content:center;width:${size}px;height:${size}px;` +
    `border-radius:${radius};background:#f4f4f4;color:#525252;font-size:${font}px;font-weight:700">` +
    `${esc(monogram(name))}</div>`
  );
}

export type Chip = { text: string; muted?: boolean };

export function chipMarkup(chip: Chip): string {
  const style = chip.muted
    ? 'display:flex;align-items:center;color:#a3a3a3;font-size:24px;font-weight:500;padding:8px 4px'
    : 'display:flex;align-items:center;background:#f4f4f4;color:#171717;font-size:24px;font-weight:500;border-radius:999px;padding:9px 18px';
  return `<div style="${style}">${esc(chip.text)}</div>`;
}

/** The shared bottom row: the hireme mark + wordmark on the left, the site
 *  domain on the right. Every card closes with this. */
export function brandFooter(): string {
  return `
  <div style="display:flex;align-items:center;justify-content:space-between">
    <div style="display:flex;align-items:center;gap:14px">
      <img src="${MARK_DATA_URI}" style="width:34px;height:34px" />
      <div style="display:flex;font-size:30px;font-weight:700;letter-spacing:-0.03em">HireMe</div>
    </div>
    <div style="display:flex;font-size:22px;color:#a3a3a3">hireme.example.com</div>
  </div>`;
}
