import { ChevronsLeft, ChevronsRight } from 'lucide-react';
import clsx from 'clsx';

import { Link } from '@@/Link';

import { useSidebarState } from './useSidebarState';
import styles from './Header.module.css';

interface Props {
  logo?: string;
}

export function Header({ logo: customLogo }: Props) {
  const { toggle, isOpen } = useSidebarState();

  return (
    <div className="flex">
      <div
        className={clsx('flex w-full flex-wrap pr-5', {
          'justify-center': !isOpen,
        })}
      >
        <Link
          to="opendocking.home"
          data-cy="opendockingSidebar-homeImage"
          className="text-2xl text-white no-underline hover:text-white hover:no-underline focus:text-white focus:no-underline focus:outline-none"
        >
          <Logo customLogo={customLogo} isOpen={isOpen} />
        </Link>
        {isOpen && (
          <div
            className={clsx(
              'space-x-1 pt-3 text-[9.4px] uppercase tracking-[.28em]',
              'text-gray-3',
              'th-dark:text-gray-warm-6'
            )}
          >
            <span className="font-medium">OpenDocking</span>
          </div>
        )}
      </div>

      <button
        type="button"
        onClick={() => toggle()}
        className={clsx(
          styles.collapseBtn,
          'flex h-6 w-6 items-center justify-center rounded border-0',
          'transition-all duration-200',
          'text-sm text-gray-4',
          'bg-graphite-900 hover:bg-graphite-500',
          'absolute',
          { '-right-[10px]': !isOpen, 'right-6': isOpen }
        )}
        aria-label="Toggle Sidebar"
        title="Toggle Sidebar"
      >
        {isOpen ? <ChevronsLeft /> : <ChevronsRight />}
      </button>
    </div>
  );
}

function getLogo(customLogo?: string) {
  if (customLogo) {
    return customLogo;
  }

  return '';
}

function Logo({
  customLogo,
  isOpen,
}: {
  customLogo?: string;
  isOpen: boolean;
}) {
  const logo = getLogo(customLogo);

  if (!logo) {
    return (
      <span className="font-bold text-white">
        {isOpen ? 'OpenDocking' : 'OD'}
      </span>
    );
  }

  return (
    <img
      src={logo}
      className={clsx('img-responsive', styles.logo, {
        '!max-h-[27px]': !isOpen,
      })}
      alt="Logo"
    />
  );
}
