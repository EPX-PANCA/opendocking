import { Loader2 } from 'lucide-react';
import { ReactNode } from 'react';

import { cn } from '@/lib/utils';

type Size = 'xs' | 'sm' | 'md';

export type Props = {
  children: ReactNode;
  className?: string;
  size?: Size;
};

const sizeStyles: Record<Size, string> = {
  xs: 'text-xs gap-1',
  sm: 'text-sm gap-2',
  md: 'text-md gap-2',
};

export function InlineLoader({ children, className, size = 'sm' }: Props) {
  return (
    <div
      className={cn(
        'flex items-center text-muted-foreground',
        sizeStyles[size],
        className
      )}
    >
      <Loader2 className="size-4 flex-none animate-spin-slow" />
      {children}
    </div>
  );
}
