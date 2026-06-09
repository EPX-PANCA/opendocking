import { cva, type VariantProps } from 'class-variance-authority';
import { PropsWithChildren } from 'react';

import { cn } from '@/lib/utils';
import { AutomationTestingProps } from '@/types';

import { Icon, IconProps } from '@@/Icon';

const badge = cva(
  [
    'inline-flex w-fit items-center gap-1 rounded-full border transition-colors',
    'focus:outline-none focus:ring-2 focus:ring-ring focus:ring-offset-2',
    '[&>a]:hover:text-inherit [&>a]:text-inherit',
  ],
  {
    variants: {
      type: {
        success:
          'border-transparent bg-success-2 text-success-9 th-dark:bg-success-10 th-dark:text-success-3 th-highcontrast:bg-success-10 th-highcontrast:text-success-3',
        warn: 'border-transparent bg-warning-2 text-warning-9 th-dark:bg-warning-10 th-dark:text-warning-3 th-highcontrast:bg-warning-10 th-highcontrast:text-warning-3',
        danger:
          'border-transparent bg-error-2 text-error-9 th-dark:bg-error-10 th-dark:text-error-3 th-highcontrast:bg-error-10 th-highcontrast:text-error-3',
        info: 'border-transparent bg-blue-2 text-blue-9 th-dark:bg-blue-10 th-dark:text-blue-3 th-highcontrast:bg-blue-10 th-highcontrast:text-blue-3',
        successSecondary:
          'border-transparent bg-success-3 text-success-9 th-dark:bg-success-9 th-dark:text-success-3 th-highcontrast:bg-success-9 th-highcontrast:text-success-3',
        warnSecondary:
          'border-transparent bg-warning-3 text-warning-9 th-dark:bg-warning-9 th-dark:text-warning-3 th-highcontrast:bg-warning-9 th-highcontrast:text-warning-3',
        dangerSecondary:
          'border-transparent bg-error-3 text-error-9 th-dark:bg-error-9 th-dark:text-error-3 th-highcontrast:bg-error-9 th-highcontrast:text-error-3',
        infoSecondary:
          'border-transparent bg-blue-3 text-blue-9 th-dark:bg-blue-9 th-dark:text-blue-3 th-highcontrast:bg-blue-9 th-highcontrast:text-blue-3',
        muted:
          'border-transparent bg-gray-3 text-gray-9 th-dark:bg-gray-9 th-dark:text-gray-3 th-highcontrast:bg-gray-9 th-highcontrast:text-gray-3',
        accent:
          'border-transparent bg-indigo-2 text-indigo-9 th-dark:bg-indigo-10 th-dark:text-indigo-3 th-highcontrast:bg-indigo-10 th-highcontrast:text-indigo-3',
        custom: '',
      },
      shape: {
        pill: '',
        rect: 'rounded',
      },
      size: {
        sm: 'px-1.5 py-px text-[10px] font-medium',
        md: 'px-2 py-0.5 text-xs font-semibold',
      },
      bordered: {
        true: 'border-solid th-highcontrast:border-white',
      },
    },
    compoundVariants: [
      {
        type: ['success', 'successSecondary'],
        bordered: true,
        className: 'border-success-4 th-dark:border-success-8',
      },
      {
        type: ['warn', 'warnSecondary'],
        bordered: true,
        className: 'border-warning-4 th-dark:border-warning-8',
      },
      {
        type: ['danger', 'dangerSecondary'],
        bordered: true,
        className: 'border-error-4 th-dark:border-error-8',
      },
      {
        type: ['info', 'infoSecondary'],
        bordered: true,
        className: 'border-blue-4 th-dark:border-blue-8',
      },
      {
        type: 'muted',
        bordered: true,
        className: 'border-gray-4 th-dark:border-gray-8',
      },
      {
        type: 'accent',
        bordered: true,
        className: 'border-indigo-4 th-dark:border-indigo-8',
      },
    ],
    defaultVariants: {
      type: 'info',
      shape: 'pill',
      size: 'md',
    },
  }
);

export type BadgeType = NonNullable<
  Exclude<VariantProps<typeof badge>['type'], 'custom'>
>;
export type BadgeShape = NonNullable<VariantProps<typeof badge>['shape']>;
export type BadgeSize = NonNullable<VariantProps<typeof badge>['size']>;

export type Props = {
  shape?: BadgeShape;
  size?: BadgeSize;
  bordered?: boolean;
  icon?: IconProps['icon'];
} & (
  | { type?: BadgeType; className?: string }
  | { type: 'custom'; className: string }
);

export function Badge({
  type = 'info',
  shape,
  size,
  className,
  bordered,
  icon,
  children,
  'data-cy': dataCy,
}: PropsWithChildren<Props> & Partial<AutomationTestingProps>) {
  return (
    <span
      className={cn(badge({ type, shape, size, bordered, className }))}
      role="status"
      data-cy={dataCy}
    >
      {icon && <Icon icon={icon} />}
      {children}
    </span>
  );
}
