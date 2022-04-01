# -*- coding: utf-8 -*-
import os
import sys
from datetime import date

from sphinx_scylladb_theme.utils import multiversion_regex_builder

# -- Global variables

# Build documentation for the following tags and branches
TAGS = []
BRANCHES = ['master', 'branch-2.2', 'branch-2.3', 'branch-2.4', 'branch-2.5', 'branch-2.6']
# Set the latest version.
LATEST_VERSION = 'branch-2.6'
# Set which versions are not released yet.
UNSTABLE_VERSIONS = ['master']
# Set which versions are deprecated
DEPRECATED_VERSIONS = ['']

# -- General configuration ------------------------------------------------

# Add any Sphinx extension module names here, as strings. They can be
# extensions coming with Sphinx (named 'sphinx.ext.*') or your custom
# ones.
extensions = [
    "sphinx.ext.autodoc",
    "sphinx.ext.todo",
    "sphinx.ext.mathjax",
    "sphinx.ext.githubpages",
    "sphinx.ext.extlinks",
    "sphinx_sitemap",
    "sphinx_scylladb_theme",
    'sphinxcontrib.datatemplates',
    "sphinx_multiversion",
    "recommonmark",
]

# The suffix(es) of source filenames.
source_suffix = [".rst", ".md"]

# The master toctree document.
master_doc = 'index'

# General information about the project.
project = 'Scylla Manager'
copyright = str(date.today().year) + ', ScyllaDB. All rights reserved.'
author = u'Scylla Project Contributors'


# List of patterns, relative to source directory, that match files and
# directories to ignore when looking for source files.
exclude_patterns = ["_build", "Thumbs.db", ".DS_Store", "_templates"]

# The name of the Pygments (syntax highlighting) style to use.
pygments_style = 'sphinx'

# -- Options for not found extension -------------------------------------------

# Template used to render the 404.html generated by this extension.
notfound_template =  '404.html'

# Prefix added to all the URLs generated in the 404 page.
notfound_urls_prefix = ''

# -- Options for redirect extension ---------------------------------------

# Read a YAML dictionary of redirections and generate an HTML file for each
redirects_file = "_utils/redirections.yaml"

# -- Options for multiversion --------------------------------------------

# Whitelist pattern for tags
smv_tag_whitelist = multiversion_regex_builder(TAGS)
# Whitelist pattern for branches
smv_branch_whitelist = multiversion_regex_builder(BRANCHES)
# Defines which version is considered to be the latest stable version.
smv_latest_version = LATEST_VERSION
# Defines the new name for the latest version.
smv_rename_latest_version = 'stable'
# Whitelist pattern for remotes (set to None to use local branches only)
smv_remote_whitelist = r"^origin$"
# Pattern for released versions
smv_released_pattern = r'^tags/.*$'
# Format for versioned output directories inside the build directory
smv_outputdir_format = '{ref.name}'

# -- Options for sitemap extension ---------------------------------------

sitemap_url_scheme = 'stable/{link}'

# -- Options for HTML output ---------------------------------------------

# The theme to use for HTML and HTML Help pages.  See the documentation for
# a list of builtin themes.
#
html_theme = 'sphinx_scylladb_theme'
templates_path = ['_templates', ]

# Theme options are theme-specific and customize the look and feel of a theme
# further.  For a list of options available for each theme, see the
# documentation.
#
html_theme_options = {
    'branch_substring_removed': 'branch-',
    'default_branch': 'master',
    'conf_py_path': 'docs/source/',
    'github_issues_repository': 'scylladb/scylla-manager',
    'github_repository': 'scylladb/scylla-manager',
    'hide_edit_this_page_button': 'false',
    'hide_version_dropdown': ['master'],
    'tag_substring_removed': 'scylla-manager-',
    'versions_unstable': UNSTABLE_VERSIONS,
    'versions_deprecated': DEPRECATED_VERSIONS,
}

# If not None, a 'Last updated on:' timestamp is inserted at every page
# bottom, using the given strftime format.
# The empty string is equivalent to '%b %d, %Y'.
#
html_last_updated_fmt = '%d %B %Y'

# Custom sidebar templates, maps document names to template names.
#
html_sidebars = {'**': ['side-nav.html']}

# Output file base name for HTML help builder.
htmlhelp_basename = 'ScyllaDocumentationdoc'

# URL which points to the root of the HTML documentation. 
html_baseurl = 'https://manager.docs.scylladb.com'

# Dictionary of values to pass into the template engine’s context for all pages
html_context = {'html_baseurl': html_baseurl}
