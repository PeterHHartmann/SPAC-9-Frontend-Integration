import { Breadcrumb, BreadcrumbItem, BreadcrumbLink, BreadcrumbList, BreadcrumbPage, BreadcrumbSeparator } from '@/components/ui/breadcrumb';
import { Separator } from '@/components/ui/separator';
import { SidebarTrigger } from '@/components/ui/sidebar';
import { capitalize } from '@/lib/utils';
import { useMemo, type PropsWithChildren } from 'react';
import type React from 'react';
import { Link, useLocation } from 'react-router-dom';

export const PageLayout: React.FC<PropsWithChildren> = ({ children }) => {
	const location = useLocation();

	const crumbs: string[] = useMemo(() => {
		return location.pathname.slice(1).split('/');
	}, [location.pathname]);

	const crumbLinks: string[] = crumbs.reduce<string[]>(
		(prev, currentValue, currentIndex) => {
			const newArr = prev.slice();
			if (currentIndex > 0) {
				newArr.push(prev[currentIndex - 1] + "/" + currentValue);
			} else {
				newArr.push(currentValue);
			}
			return newArr;
		},
		[]
	);

	return (
		<>
			<header className="flex h-16 shrink-0 items-center gap-2 border-b px-4">
				<SidebarTrigger className="-ml-1" />
				<Separator orientation="vertical" className="mr-2 h-4" />
				<Breadcrumb>
					<BreadcrumbList>
						<BreadcrumbItem className="hidden md:block">
							<BreadcrumbLink href="/">
								Movie Quotes
							</BreadcrumbLink>
						</BreadcrumbItem >
						{crumbs.map((crumb, index) => (
							<>
								<BreadcrumbSeparator className="hidden md:block" />
								<BreadcrumbItem className="hidden md:block">
									{index < crumbs.length - 1
										?
										<Link to={crumbLinks[index]}>
											{capitalize(crumb)}
										</Link>
										:
										<BreadcrumbPage>{capitalize(crumb)}</BreadcrumbPage>
									}
								</BreadcrumbItem>
							</>
						))
						}
					</BreadcrumbList>
				</Breadcrumb>
			</header>
			<main className="flex flex-1 w-full p-8 bg-white min-h-[calc(100vh-64px)]">
				{children}
			</main>
		</>
	);
};