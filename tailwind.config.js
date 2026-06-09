const plugin = require('tailwindcss/plugin');
const defaultTheme = require('tailwindcss/defaultTheme');
const colors = require('./app/assets/css/colors.json');

/** @type {import('tailwindcss').Config} */
module.exports = {
    darkMode: ['class'],
    content: ['./app/**/*.{html,tsx}'],
  corePlugins: {
    preflight: false,
  },
  theme: {
  	colors: {
  		transparent: 'transparent',
  		current: 'currentColor',
  		inherit: 'inherit',
            ...colors,
  		'group-accent': 'colors.violet',
  		'legacy-grey-3': 'var(--grey-3)',
  		'legacy-blue-2': 'var(--blue-2)',
  		'legacy-blue-9': 'var(--blue-9)',
  		border: 'var(--border)',
  		input: 'var(--input)',
  		ring: 'var(--ring)',
  		background: 'var(--background)',
  		foreground: 'var(--foreground)',
  		primary: {
  			DEFAULT: 'var(--primary)',
  			foreground: 'var(--primary-foreground)'
  		},
  		secondary: {
  			DEFAULT: 'var(--secondary)',
  			foreground: 'var(--secondary-foreground)'
  		},
  		destructive: {
  			DEFAULT: 'var(--destructive)',
  			foreground: 'var(--destructive-foreground)'
  		},
  		muted: {
  			DEFAULT: 'var(--muted)',
  			foreground: 'var(--muted-foreground)'
  		},
  		accent: {
  			DEFAULT: 'var(--accent)',
  			foreground: 'var(--accent-foreground)'
  		},
  		popover: {
  			DEFAULT: 'var(--popover)',
  			foreground: 'var(--popover-foreground)'
  		},
  		card: {
  			DEFAULT: 'var(--card)',
  			foreground: 'var(--card-foreground)'
  		},
  		sidebar: {
  			DEFAULT: 'var(--sidebar-background)',
  			foreground: 'var(--sidebar-foreground)',
  			primary: 'var(--sidebar-primary)',
  			'primary-foreground': 'var(--sidebar-primary-foreground)',
  			accent: 'var(--sidebar-accent)',
  			'accent-foreground': 'var(--sidebar-accent-foreground)',
  			border: 'var(--sidebar-border)',
  			ring: 'var(--sidebar-ring)'
  		}
  	},
  	extend: {
  		fontFamily: {
  			sans: [
  				'Inter',
                    ...defaultTheme.fontFamily.sans
                ]
  		},
  		animation: {
  			'spin-slow': 'spin 2s linear infinite',
  			'accordion-down': 'accordion-down 0.2s ease-out',
  			'accordion-up': 'accordion-up 0.2s ease-out'
  		},
  		keyframes: {
  			'accordion-down': {
  				from: {
  					height: '0'
  				},
  				to: {
  					height: 'var(--radix-accordion-content-height)'
  				}
  			},
  			'accordion-up': {
  				from: {
  					height: 'var(--radix-accordion-content-height)'
  				},
  				to: {
  					height: '0'
  				}
  			}
  		}
  	}
  },

  plugins: [
    plugin(({ addVariant }) => {
      addVariant('be', '&:is([data-edition="BE"] *)');
      addVariant('th-highcontrast', '&:is([theme="highcontrast"] *)');
      addVariant('th-dark', '&:is([theme="dark"] *)');
    }),
    plugin(function ({ addVariant }) {
      addVariant('progress-filled', ['&::-webkit-progress-value', '&::-moz-progress-bar']);
    }),
    require('tailwindcss-animate'),
  ],
};
