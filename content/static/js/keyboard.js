/**
 * @license
 * Copyright 2019-2020 The Go Authors. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */

// This file implements the behavior of the keyboard shortcut which allows
// for users to:
// - press 'y' to to change browser URL to the canonical URL without triggering
// a reload.
// - press 'esc' set focus ouf of input element.
// - press 's' to set focus at search box.

const canonicalURLPath = document.querySelector(".js-canonicalURLPath") && document.querySelector(".js-canonicalURLPath").dataset[
  "canonicalUrlPath"
];

document.addEventListener("keydown", (e) => {
  // TODO(golang.org/issue/40246): consolidate keyboard shortcut behavior across the site.
  const t = e.target.tagName;
  switch (t) {
    case "INPUT":
      switch (e.key) {
        case "Esc":
        case "Escape":
          e.target.blur();
          e.preventDefault();
          return;
      }
    case "SELECT":
    case "TEXTAREA":
      return;
  }
  if (e.target.isContentEditable) {
    return;
  }
  if (e.metaKey || e.ctrlKey) {
    return;
  }
  switch (e.key) {
    case "S":
    case "s":
      document.querySelector('input[name="q"]').focus()
      e.preventDefault();
      break;
    case "y":
      if (canonicalURLPath && canonicalURLPath !== "") {
        window.history.replaceState(null, "", canonicalURLPath);
      }
      break;
  }
});
